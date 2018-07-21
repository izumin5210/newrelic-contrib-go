package nrsql

import (
	"context"
	"database/sql"
)

// Queryer represents the object that can perform SQL queries.
type Queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// PreparedQueryer represents the object that can perform prepared SQL queries.
type PreparedQueryer interface {
	QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
}

// Execer represents the object that can execute SQL queries.
type Execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// PreparedExecer represents the object that can execute prepared SQL queries.
type PreparedExecer interface {
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
}

// Preparer represents the object that can create SQL prepared statements.
type Preparer interface {
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}
