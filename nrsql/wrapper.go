package nrsql

import (
	"context"
	"database/sql"

	"github.com/izumin5210/newrelic-contrib-go/nrutil"
	newrelic "github.com/newrelic/go-agent"
)

func wrapQueryer(queryer Queryer, cfg *Config) Queryer {
	return &queryerWrapper{original: queryer, config: cfg}
}

type queryerWrapper struct {
	original Queryer
	config   *Config
}

func (w *queryerWrapper) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	segment(ctx, w.config, parseQuery(query), args, func() {
		rows, err = w.original.QueryContext(ctx, query, args...)
	})
	return
}

func (w *queryerWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) (row *sql.Row) {
	segment(ctx, w.config, parseQuery(query), args, func() {
		row = w.original.QueryRowContext(ctx, query, args...)
	})
	return
}

func wrapExecer(execer Execer, cfg *Config) Execer {
	return &execerWrapper{original: execer, config: cfg}
}

type execerWrapper struct {
	original Execer
	config   *Config
}

func (w *execerWrapper) ExecContext(ctx context.Context, query string, args ...interface{}) (res sql.Result, err error) {
	segment(ctx, w.config, parseQuery(query), args, func() {
		res, err = w.original.ExecContext(ctx, query, args...)
	})
	return
}

func wrapPreparedQueryer(queryer PreparedQueryer, query *query, cfg *Config) PreparedQueryer {
	return &preparedQueryerWrapper{original: queryer, query: query, config: cfg}
}

type preparedQueryerWrapper struct {
	original PreparedQueryer
	query    *query
	config   *Config
}

func (w *preparedQueryerWrapper) QueryContext(ctx context.Context, args ...interface{}) (rows *sql.Rows, err error) {
	segment(ctx, w.config, w.query, args, func() {
		rows, err = w.original.QueryContext(ctx, args...)
	})
	return
}

func (w *preparedQueryerWrapper) QueryRowContext(ctx context.Context, args ...interface{}) (row *sql.Row) {
	segment(ctx, w.config, w.query, args, func() {
		row = w.original.QueryRowContext(ctx, args...)
	})
	return
}

func wrapPreparedExecer(execer PreparedExecer, query *query, cfg *Config) PreparedExecer {
	return &preparedExecerWrapper{original: execer, query: query, config: cfg}
}

type preparedExecerWrapper struct {
	original PreparedExecer
	query    *query
	config   *Config
}

func (w *preparedExecerWrapper) ExecContext(ctx context.Context, args ...interface{}) (res sql.Result, err error) {
	segment(ctx, w.config, w.query, args, func() {
		res, err = w.original.ExecContext(ctx, args...)
	})
	return
}

func segment(ctx context.Context, cfg *Config, q *query, args []interface{}, do func()) {
	seg := &newrelic.DatastoreSegment{
		StartTime:          newrelic.StartSegmentNow(nrutil.Transaction(ctx)),
		Product:            cfg.Datastore,
		Collection:         q.TableName,
		Operation:          q.Operation,
		ParameterizedQuery: q.Raw,
		Host:               cfg.Host,
		PortPathOrID:       cfg.PortPathOrID,
		DatabaseName:       cfg.DBName,
	}
	defer seg.End()

	do()
}
