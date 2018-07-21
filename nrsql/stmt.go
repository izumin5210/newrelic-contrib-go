package nrsql

import "database/sql"

type Stmt interface {
	PreparedQueryer
	PreparedExecer
	Closer

	Stmt() *sql.Stmt
}

type stmtWrapper struct {
	original *sql.Stmt
	PreparedQueryer
	PreparedExecer
}

func wrapStmt(stmt *sql.Stmt) Stmt {
	return &stmtWrapper{
		original:        stmt,
		PreparedQueryer: wrapPreparedQueryer(stmt),
		PreparedExecer:  wrapPreparedExecer(stmt),
	}
}

func (w *stmtWrapper) Close() error {
	return w.original.Close()
}

func (w *stmtWrapper) Stmt() *sql.Stmt {
	return w.original
}
