package nrsql

import (
	"context"
	"database/sql"

	"github.com/izumin5210/newrelic-contrib-go/nrutil"
	newrelic "github.com/newrelic/go-agent"
)

func wrapQueryer(queryer Queryer) Queryer {
	return &queryerWrapper{original: queryer}
}

type queryerWrapper struct {
	original Queryer
}

func (w *queryerWrapper) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	segment(ctx, parseQuery(query), args, func() {
		rows, err = w.original.QueryContext(ctx, query, args...)
	})
	return
}

func (w *queryerWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) (row *sql.Row) {
	segment(ctx, parseQuery(query), args, func() {
		row = w.original.QueryRowContext(ctx, query, args...)
	})
	return
}

func wrapExecer(execer Execer) Execer {
	return &execerWrapper{original: execer}
}

type execerWrapper struct {
	original Execer
}

func (w *execerWrapper) ExecContext(ctx context.Context, query string, args ...interface{}) (res sql.Result, err error) {
	segment(ctx, parseQuery(query), args, func() {
		res, err = w.original.ExecContext(ctx, query, args...)
	})
	return
}

func wrapPreparedQueryer(queryer PreparedQueryer, query *query) PreparedQueryer {
	return &preparedQueryerWrapper{original: queryer, query: query}
}

type preparedQueryerWrapper struct {
	original PreparedQueryer
	query    *query
}

func (w *preparedQueryerWrapper) QueryContext(ctx context.Context, args ...interface{}) (rows *sql.Rows, err error) {
	segment(ctx, w.query, args, func() {
		rows, err = w.original.QueryContext(ctx, args...)
	})
	return
}

func (w *preparedQueryerWrapper) QueryRowContext(ctx context.Context, args ...interface{}) (row *sql.Row) {
	segment(ctx, w.query, args, func() {
		row = w.original.QueryRowContext(ctx, args...)
	})
	return
}

func wrapPreparedExecer(execer PreparedExecer, query *query) PreparedExecer {
	return &preparedExecerWrapper{original: execer, query: query}
}

type preparedExecerWrapper struct {
	original PreparedExecer
	query    *query
}

func (w *preparedExecerWrapper) ExecContext(ctx context.Context, args ...interface{}) (res sql.Result, err error) {
	segment(ctx, w.query, args, func() {
		res, err = w.original.ExecContext(ctx, args...)
	})
	return
}

func segment(ctx context.Context, q *query, args []interface{}, do func()) {
	seg := &newrelic.DatastoreSegment{
		StartTime: newrelic.StartSegmentNow(nrutil.Transaction(ctx)),
		// Product:            "",
		Collection:         q.TableName,
		Operation:          q.Operation,
		ParameterizedQuery: q.Raw,
		// QueryParameters:    map[string]interface{}{},
		// Host:               "",
		// PortPathOrID:       "",
		// DatabaseName:       "",
	}
	defer seg.End()

	do()
}
