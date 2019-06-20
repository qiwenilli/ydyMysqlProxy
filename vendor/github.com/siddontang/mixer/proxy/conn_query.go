package proxy

import (
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"github.com/siddontang/go-log/log"
	"github.com/siddontang/mixer/client"
	. "github.com/siddontang/mixer/mysql"

	//
	"reflect"
	"sync"
)

func (c *Conn) handleQuery(sql string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("execute %s error %v", sql, e)
			return
		}
	}()

	//qiwen
	_parser := parser.New()
	stmt, err := _parser.ParseOneStmt(sql, "", "")
	if err != nil {
		return fmt.Errorf("statement %s not support now", err)
	}

	//debug info
	t := reflect.TypeOf(stmt)
	log.Debug("exec sql: ", t, " | ", sql)

	switch v := stmt.(type) {
	case *ast.ShowStmt, *ast.SelectStmt:
		return c.handleSelect(sql)
	case *ast.InsertStmt, *ast.UpdateStmt, *ast.DeleteStmt, *ast.AlterTableStmt:
		return c.handleExec(sql)
    // case *ast.SetStmt:
	// 	return c.handleSet(sql)
	case *ast.BeginStmt:
		return c.handleBegin()
	case *ast.RollbackStmt:
		return c.handleRollback()
	case *ast.CommitStmt:
		return c.handleCommit()
	default:
		return fmt.Errorf("statement %T not support now", v)
	}
	//qiwen end
	return nil
}

func (c *Conn) getShardList() ([]*Node, error) {
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	ns := []string{"node1"}

	if len(ns) == 0 {
		return nil, nil
	}

	n := make([]*Node, 0, len(ns))
	for _, name := range ns {
		n = append(n, c.server.getNode(name))
	}
	return n, nil
}

func (c *Conn) getConn(n *Node, isSelect bool) (co *client.SqlConn, err error) {
	if !c.needBeginTx() {
		if isSelect {
			co, err = n.getSelectConn()
		} else {
			co, err = n.getMasterConn()
		}
		if err != nil {
			return
		}
	} else {
		var ok bool
		c.Lock()
		co, ok = c.txConns[n]
		c.Unlock()

		if !ok {
			if co, err = n.getMasterConn(); err != nil {
				return
			}

			if err = co.Begin(); err != nil {
				return
			}

			c.Lock()
			c.txConns[n] = co
			c.Unlock()
		}
	}

	//todo, set conn charset, etc...
	if err = co.UseDB(c.schema.db); err != nil {
		return
	}

	if err = co.SetCharset(c.charset); err != nil {
		return
	}

	return
}

func (c *Conn) getShardConns(isSelect bool) ([]*client.SqlConn, error) {
	nodes, err := c.getShardList()
	if err != nil {
		return nil, err
	} else if nodes == nil {
		return nil, nil
	}

	conns := make([]*client.SqlConn, 0, len(nodes))

	var co *client.SqlConn
	for _, n := range nodes {
		co, err = c.getConn(n, isSelect)
		if err != nil {
			break
		}

		conns = append(conns, co)
	}

	return conns, err
}

func (c *Conn) executeInShard(conns []*client.SqlConn, sql string) ([]*Result, error) {
	var wg sync.WaitGroup
	wg.Add(len(conns))

	rs := make([]interface{}, len(conns))

	f := func(rs []interface{}, i int, co *client.SqlConn) {
		r, err := co.Execute(sql)
		if err != nil {
			rs[i] = err
		} else {
			rs[i] = r
		}

		wg.Done()
	}

	for i, co := range conns {
		go f(rs, i, co)
	}

	wg.Wait()

	var err error
	r := make([]*Result, len(conns))
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*Result)
	}

	return r, err
}

func (c *Conn) closeShardConns(conns []*client.SqlConn, rollback bool) {
	if c.isInTransaction() {
		return
	}

	for _, co := range conns {
		if rollback {
			co.Rollback()
		}

		co.Close()
	}
}

func (c *Conn) newEmptyResultset() *Resultset {
	r := new(Resultset)
	r.Fields = make([]*Field, 1)
	//
	r.Values = make([][]interface{}, 0)
	r.RowDatas = make([]RowData, 0)

	return r
}

func (c *Conn) handleSelect(sql string) error {

	conns, err := c.getShardConns(true)
	if err != nil {
		return err
	} else if conns == nil {
		r := c.newEmptyResultset()
		return c.writeResultset(c.status, r)
	}

	var rs []*Result

	rs, err = c.executeInShard(conns, sql)

	c.closeShardConns(conns, false)

	if err == nil {
		err = c.mergeSelectResult(rs)
	}

	return err
}

func (c *Conn) beginShardConns(conns []*client.SqlConn) error {
	if c.isInTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Begin(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) commitShardConns(conns []*client.SqlConn) error {
	if c.isInTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) handleExec(sql string) error {

	conns, err := c.getShardConns(false)
	if err != nil {
		return err
	} else if conns == nil {
		return c.writeOK(nil)
	}

	var rs []*Result

	if len(conns) == 1 {
		rs, err = c.executeInShard(conns, sql)
	} else {
		//for multi nodes, 2PC simple, begin, exec, commit
		//if commit error, data maybe corrupt
		for {
			if err = c.beginShardConns(conns); err != nil {
				break
			}

			if rs, err = c.executeInShard(conns, sql); err != nil {
				break
			}

			err = c.commitShardConns(conns)
			break
		}
	}

	c.closeShardConns(conns, err != nil)

	if err == nil {
		err = c.mergeExecResult(rs)
	}

	return err
}

func (c *Conn) mergeExecResult(rs []*Result) error {
	r := new(Result)

	for _, v := range rs {
		r.Status |= v.Status
		r.AffectedRows += v.AffectedRows
		if r.InsertId == 0 {
			r.InsertId = v.InsertId
		} else if r.InsertId > v.InsertId {
			//last insert id is first gen id for multi row inserted
			//see http://dev.mysql.com/doc/refman/5.6/en/information-functions.html#function_last-insert-id
			r.InsertId = v.InsertId
		}
	}

	if r.InsertId > 0 {
		c.lastInsertId = int64(r.InsertId)
	}

	c.affectedRows = int64(r.AffectedRows)

	return c.writeOK(r)
}

func (c *Conn) mergeSelectResult(rs []*Result) error {
	r := rs[0].Resultset

	status := c.status | rs[0].Status

	for i := 1; i < len(rs); i++ {
		status |= rs[i].Status

		//check fields equal

		for j := range rs[i].Values {
			r.Values = append(r.Values, rs[i].Values[j])
			r.RowDatas = append(r.RowDatas, rs[i].RowDatas[j])
		}
	}

	//to do order by, group by, limit offset
	//to do, add log here, sort may error because order by key not exist in resultset fields

	return c.writeResultset(status, r)
}
