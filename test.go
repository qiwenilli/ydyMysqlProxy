package main

import (
	"fmt"
	"github.com/vitessio/vitess/go/vt/sqlparser"
    "github.com/qiwenilli/sqlfilter"
)

type SqlFilter struct{
    PrefixTableName string
}

func main() {

	sqla := []string{
		"SELECT ku11111.t.a from ku11111.t limit 10",
		"SELECT ku11111.t.a,(select ku11111.t.a from ku11111.t) from (select * from ku11111.t) b",
		"SELECT ku11111.t.a,max(ku11111.t.b) from ku11111.t limit 10",
		"SELECT  a111.employees.a,max(a11.employees.b) FROM a111.employees" +
			" UNION " +
			"SELECT  b111.employees.a FROM b111.employees" +
			" UNION " +
			"SELECT  c111.employees.a FROM c111.employees",
		"SELECT ku11111.t.a from ku11111.t left join ku11111.t1 on ku11111.t.a=ku11111.t.b and ku11111.t.a=ku11111.t.b or ku11111.t.a=ku11111.t.b limit 10",
		"SELECT ku11111.t.a from ku11111.t right join ku11111.t1 on ku11111.t.a=ku11111.t.b and ku11111.t.a=ku11111.t.b or ku11111.t.a=ku11111.t.b limit 10",
		"SELECT ku11111.t.a from ku11111.t inner join ku11111.t1 on ku11111.t.a=ku11111.t.b and ku11111.t.a=ku11111.t.b or ku11111.t.a=ku11111.t.b limit 10",
		"SELECT ku11111.t.a from ku11111.t where ku11111.t.a=1 and ku11111.t.a=1 or ku11111.t.a=1 and ku11111.t.a not in (ku11111.t.a,ku11111.t.b) limit 10",
		"SELECT ku11111.t.a,(select ku11111.t.a from ku11111.t.a)  from ku11111.t where ku11111.t.a=1 and ku11111.t.a=1 or ku11111.t.a=1 and ku11111.t.a not in (ku11111.t.a,ku11111.t.b) limit 10",
		"SELECT ku11111.t.a,(select ku11111.t.a from ku11111.t)  from ku11111.t",
		"SELECT * from ku11111.t group by ku11111.t.a,1 having ku11111.t.a order by ku11111.t.a",
		"SELECT distinct(select ku11111.t.a from ku11111.t) from ku11111.t group by ku11111.t.a,1 having ku11111.t.a order by ku11111.t.a",
        "select * from a.a",
	}

	//

	for _, sql := range sqla {
		//
		stmt, err := sqlparser.Parse(sql)
		if err != nil {
			fmt.Println(err)
			continue
		}

        _sqlFilter := sqlfilter.Filter{
            PrefixTableName:"aaaaa_",
        }
		_sqlFilter.SqlMain(stmt)

		//
		buf1 := sqlparser.NewTrackedBuffer(nil)
		stmt.Format(buf1)

		fmt.Println("end: ", buf1.String())

	}
}

