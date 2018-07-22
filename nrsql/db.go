package nrsql

import (
	"context"
	"database/sql"
)

// DB wraps a *sql.DB object.
type DB interface {
	Queryer
	Execer
	Preparer
	Pinger
	Beginner
	Closer

	DB() *sql.DB
}

type dbWrapper struct {
	original *sql.DB
	Queryer
	Execer

	config *Config
}

// Wrap wraps a *sql.DB object to measure performances and sent them to New Relic.
func Wrap(db *sql.DB, opts ...Option) DB {
	cfg := createConfig(opts)
	return &dbWrapper{
		original: db,
		Queryer:  wrapQueryer(db, cfg),
		Execer:   wrapExecer(db, cfg),
		config:   cfg,
	}
}

func (w *dbWrapper) Prepare(query string) (Stmt, error) {
	return w.PrepareContext(context.Background(), query)
}

func (w *dbWrapper) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	stmt, err := w.original.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return wrapStmt(stmt, w.config, parseQuery(query)), nil
}

func (w *dbWrapper) Ping() error {
	return w.original.Ping()
}

func (w *dbWrapper) PingContext(ctx context.Context) error {
	return w.original.PingContext(ctx)
}

func (w *dbWrapper) Close() error {
	return w.original.Close()
}

func (w *dbWrapper) Begin() (Tx, error) {
	tx, err := w.original.Begin()
	if err != nil {
		return nil, err
	}
	return wrapTx(tx, w.config), nil
}

func (w *dbWrapper) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	tx, err := w.original.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return wrapTx(tx, w.config), nil
}

func (w *dbWrapper) DB() *sql.DB {
	return w.original
}
