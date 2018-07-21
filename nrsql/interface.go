package nrsql

import (
	"context"
	"database/sql"
)

type Queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type PreparedQueryer interface {
	QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type PreparedExecer interface {
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
}

type Preparer interface {
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

type Closer interface {
	Close() error
}
