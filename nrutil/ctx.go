package nrutil

import (
	"context"

	"github.com/newrelic/go-agent"
)

type key struct{}

var (
	txnKey key
)

// Transaction extracts newrelic transaction object from request context
func Transaction(ctx context.Context) newrelic.Transaction {
	v := ctx.Value(txnKey)
	if v == nil {
		return nil
	}
	if txn, ok := v.(newrelic.Transaction); ok {
		return txn
	}
	return nil
}

// SetTransaction stores newrelic transaction object into given context.
func SetTransaction(ctx context.Context, txn newrelic.Transaction) context.Context {
	return context.WithValue(ctx, txnKey, txn)
}
