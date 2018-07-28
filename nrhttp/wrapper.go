package nrhttp

import (
	"net/http"

	"github.com/izumin5210/newrelic-contrib-go/nrutil"
	newrelic "github.com/newrelic/go-agent"
)

// WrapHandle wraps http.Handler and set newrelic.Transaction into http.Request object.
func WrapHandle(app newrelic.Application, pattern string, handler http.Handler) (string, http.Handler) {
	return newrelic.WrapHandle(app, pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, injectTxnToRequest(w, r))
	}))
}

// WrapHandleFunc wraps http.Handler and set newrelic.Transaction into http.Request object.
func WrapHandleFunc(app newrelic.Application, pattern string, handler func(http.ResponseWriter, *http.Request)) (string, func(http.ResponseWriter, *http.Request)) {
	return newrelic.WrapHandleFunc(app, pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(w, injectTxnToRequest(w, r))
	})
}

func injectTxnToRequest(w http.ResponseWriter, r *http.Request) *http.Request {
	if txn, ok := w.(newrelic.Transaction); ok {
		r = r.WithContext(nrutil.SetTransaction(r.Context(), txn))
	}
	return r
}

// NewRoundTripper creates an http.RoundTripper to measure external requests.
func NewRoundTripper(original http.RoundTripper) http.RoundTripper {
	if original == nil {
		original = http.DefaultTransport
	}
	return &roundTripperImpl{original: original}
}

type roundTripperImpl struct {
	original http.RoundTripper
}

func (rt *roundTripperImpl) RoundTrip(req *http.Request) (*http.Response, error) {
	txn := nrutil.Transaction(req.Context())
	seg := newrelic.StartExternalSegment(txn, req)

	resp, err := rt.original.RoundTrip(req)

	seg.Response = resp
	seg.End()

	return resp, err
}
