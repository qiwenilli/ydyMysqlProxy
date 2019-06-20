package sqlfilter

import (
	"fmt"
	"github.com/vitessio/vitess/go/vt/sqlparser"
	"reflect"
	// "github.com/go-ffmt/ffmt"
)

type Filter struct {
	PrefixTableName string
}

func (this *Filter) SqlMain(stmt interface{}) interface{} {

	switch v := stmt.(type) {
	case *sqlparser.Subquery:
		this.SqlMain(v.Select)
	case *sqlparser.Select:
		this.sqlSelect(v)
	case *sqlparser.Union:
		this.sqlUnion(v)
	case *sqlparser.AliasedExpr:
		this.SqlMain(v.Expr)
	case *sqlparser.ColName:
		v.Qualifier = this.sqlTableName(v.Qualifier)
	case sqlparser.TableName:
		//不是指针，直接返回
		stmt = this.sqlTableName(v)
	case *sqlparser.FuncExpr:
		this.sqlFunc(v)
	case *sqlparser.AliasedTableExpr:
		rst := this.SqlMain(v.Expr)
		//
		switch rst.(type) {
		case sqlparser.TableName:
			//不是指针，直接返回,需要重新赋值
			v.Expr = rst.(sqlparser.TableName)
		}
	case *sqlparser.JoinTableExpr:
		this.SqlMain(v.LeftExpr)
		this.SqlMain(v.RightExpr)
		this.SqlMain(v.Condition)
	case sqlparser.JoinCondition:
		rst := this.SqlMain(v.On)
		//
		switch rst.(type) {
		case sqlparser.JoinCondition:
			//不是指针，直接返回,需要重新赋值
			stmt = rst.(sqlparser.JoinCondition)
		}
	case *sqlparser.ComparisonExpr:
		this.SqlMain(v.Left)
		this.SqlMain(v.Right)
	case *sqlparser.OrExpr:
		this.SqlMain(v.Left)
		this.SqlMain(v.Right)
	case *sqlparser.AndExpr:
		this.SqlMain(v.Left)
		this.SqlMain(v.Right)
	case *sqlparser.Where:
		this.SqlMain(v.Expr)
	case sqlparser.Expr:
		this.sqlExpr(v)
	case *sqlparser.SQLVal:
		println("val")
	default:
		fmt.Println(reflect.TypeOf(v))
	}
	return stmt
}

func (this *Filter) sqlUnion(stmt *sqlparser.Union) {
	this.SqlMain(stmt.Left)
	this.SqlMain(stmt.Right)
}

func (this *Filter) sqlSelect(v *sqlparser.Select) {

	for _, vv := range v.SelectExprs {
		this.SqlMain(vv)
	}

	for _, vv := range v.From {
		this.SqlMain(vv)
	}

	if v.Where != nil {
		this.SqlMain(v.Where)
	}

	if v.GroupBy != nil {
		for _, vv := range v.GroupBy {
			this.SqlMain(vv)
		}
	}

	if v.Having != nil {
		this.SqlMain(v.Having)
	}

	if len(v.OrderBy) > 0 {
		for _, vv := range v.OrderBy {
			this.SqlMain(vv.Expr)
		}
	}
}

func (this *Filter) sqlFunc(e *sqlparser.FuncExpr) {
	for _, v := range e.Exprs {
		this.SqlMain(v)
	}
}

func (this *Filter) sqlExpr(stmt sqlparser.Expr) {
	switch v := stmt.(type) {
	case sqlparser.ValTuple:
		for _, vv := range v {
			this.SqlMain(vv)
		}
	}
}

func (this *Filter) sqlTableName(t sqlparser.TableName) sqlparser.TableName {
	if t.Qualifier.String() != "" {
		return sqlparser.TableName{Name: t.Name, Qualifier: sqlparser.NewTableIdent(this.PrefixTableName + t.Qualifier.String())}
	}

	// 可用于字段过滤
	// if c.Name.String() == "abc" {
	//
	// 	_colident := sqlparser.NewColIdent("md5(abcd)")
	//
	// 	c.Name = _colident
	// }
	return t
}
