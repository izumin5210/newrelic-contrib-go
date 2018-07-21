package nrsql

import (
	"context"
	"database/sql"
)

func wrapQueryer(queryer Queryer) Queryer {
	return &queryerWrapper{original: queryer}
}

type queryerWrapper struct {
	original Queryer
}

func (w *queryerWrapper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// TODO
	return w.original.QueryContext(ctx, query, args...)
}

func (w *queryerWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	// TODO
	return w.original.QueryRowContext(ctx, query, args...)
}

func wrapExecer(execer Execer) Execer {
	return &execerWrapper{original: execer}
}

type execerWrapper struct {
	original Execer
}

func (w *execerWrapper) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// TODO
	return w.original.ExecContext(ctx, query, args...)
}

func wrapPreparedQueryer(queryer PreparedQueryer) PreparedQueryer {
	return &preparedQueryerWrapper{original: queryer}
}

type preparedQueryerWrapper struct {
	original PreparedQueryer
}

func (w *preparedQueryerWrapper) QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	// TODO
	return w.original.QueryContext(ctx, args...)
}

func (w *preparedQueryerWrapper) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row {
	// TODO
	return w.original.QueryRowContext(ctx, args...)
}

func wrapPreparedExecer(execer PreparedExecer) PreparedExecer {
	return &preparedExecerWrapper{original: execer}
}

type preparedExecerWrapper struct {
	original PreparedExecer
}

func (w *preparedExecerWrapper) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	// TODO
	return w.original.ExecContext(ctx, args...)
}
