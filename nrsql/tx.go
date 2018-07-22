package nrsql

import (
	"context"
	"database/sql"
)

// Tx wraps a *sql.Tx object.
type Tx interface {
	Queryer
	Execer
	Transactor

	StmtContext(context.Context, Stmt) Stmt

	Tx() *sql.Tx
}

type txWrapper struct {
	original *sql.Tx
	Queryer
	Execer

	config *Config
}

func wrapTx(tx *sql.Tx, cfg *Config) Tx {
	return &txWrapper{
		original: tx,
		Queryer:  wrapQueryer(tx, cfg),
		Execer:   wrapExecer(tx, cfg),
		config:   cfg,
	}
}

func (w *txWrapper) StmtContext(ctx context.Context, stmt Stmt) Stmt {
	return wrapStmt(w.original.StmtContext(ctx, stmt.Stmt()), w.config, stmt.parsedQuery())
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
