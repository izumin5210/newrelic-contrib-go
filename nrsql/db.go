package nrsql

import (
	"context"
	"database/sql"
)

type DB interface {
	Queryer
	Execer
	Preparer
	Closer

	Begin() (Tx, error)
	BeginTx(context.Context, *sql.TxOptions) (Tx, error)

	DB() *sql.DB
}

type dbWrapper struct {
	original *sql.DB
	Queryer
	Execer
}

func Wrap(db *sql.DB) DB {
	return &dbWrapper{
		original: db,
		Queryer:  wrapQueryer(db),
		Execer:   Execer(db),
	}
}

func (w *dbWrapper) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	stmt, err := w.original.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return wrapStmt(stmt), nil
}

func (w *dbWrapper) Close() error {
	return w.original.Close()
}

func (w *dbWrapper) Begin() (Tx, error) {
	tx, err := w.original.Begin()
	if err != nil {
		return nil, err
	}
	return wrapTx(tx), nil
}

func (w *dbWrapper) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	tx, err := w.original.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return wrapTx(tx), nil
}

func (w *dbWrapper) DB() *sql.DB {
	return w.original
}
