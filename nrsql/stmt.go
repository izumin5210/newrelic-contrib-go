package nrsql

import "database/sql"

// Stmt wraps a *sql.Stmt object.
type Stmt interface {
	PreparedQueryer
	PreparedExecer

	Stmt() *sql.Stmt
}

type stmtWrapper struct {
	original *sql.Stmt
	PreparedQueryer
	PreparedExecer
}

func wrapStmt(stmt *sql.Stmt, query *query) Stmt {
	return &stmtWrapper{
		original:        stmt,
		PreparedQueryer: wrapPreparedQueryer(stmt, query),
		PreparedExecer:  wrapPreparedExecer(stmt, query),
	}
}

func (w *stmtWrapper) Close() error {
	return w.original.Close()
}

func (w *stmtWrapper) Stmt() *sql.Stmt {
	return w.original
}
