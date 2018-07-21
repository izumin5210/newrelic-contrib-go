package nrsql

import (
	"strings"

	"github.com/xwb1989/sqlparser"
)

type query struct {
	Operation string
	TableName string
	Raw       string
}

func parseQuery(queryStr string) *query {
	q := &query{Raw: queryStr}

	stmt, err := sqlparser.Parse(queryStr)
	if err != nil {
		q.Operation = strings.Split(queryStr, " ")[0]
		return q
	}

	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		q.Operation = "SELECT"
		q.TableName = getTableNameFrom(stmt.From).CompliantName()
	case *sqlparser.Insert:
		q.Operation = "INSERT"
		q.TableName = stmt.Table.Name.CompliantName()
	case *sqlparser.Update:
		q.Operation = "UPDATE"
		q.TableName = getTableNameFrom(stmt.TableExprs).CompliantName()
	case *sqlparser.Delete:
		q.Operation = "DELETE"
		q.TableName = getTableNameFrom(stmt.TableExprs).CompliantName()
	default:
		q.Operation = strings.Split(queryStr, " ")[0]
	}

	return q
}

func getTableNameFrom(exprs sqlparser.TableExprs) sqlparser.TableIdent {
	for _, expr := range exprs {
		if tbl, ok := expr.(*sqlparser.AliasedTableExpr); ok {
			return sqlparser.GetTableName(tbl.Expr)
		}
	}
	return sqlparser.NewTableIdent("")
}
