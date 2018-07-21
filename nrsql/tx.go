package nrsql

import (
	"context"
	"database/sql"
)

type Tx interface {
	Queryer
	Execer

	StmtContext(context.Context, Stmt) Stmt
	Commit() error
	Rollback() error

	Tx() *sql.Tx
}

type txWrapper struct {
	original *sql.Tx
	Queryer
	Execer
}

func wrapTx(tx *sql.Tx) Tx {
	return &txWrapper{
		original: tx,
		Queryer:  wrapQueryer(tx),
		Execer:   wrapExecer(tx),
	}
}

func (w *txWrapper) StmtContext(ctx context.Context, stmt Stmt) Stmt {
	return wrapStmt(w.original.StmtContext(ctx, stmt.Stmt()))
}

func (w *txWrapper) Commit() error {
	return w.original.Commit()
}

func (w *txWrapper) Rollback() error {
	return w.original.Rollback()
}

func (w *txWrapper) Tx() *sql.Tx {
	return w.original
}
