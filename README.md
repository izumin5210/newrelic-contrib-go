# newrelic-contrib for Golang
[![Build Status](https://travis-ci.com/izumin5210/newrelic-contrib-go.svg?branch=master)](https://travis-ci.com/izumin5210/newrelic-contrib-go)
[![GoDoc](https://godoc.org/github.com/izumin5210/newrelic-contrib-go?status.svg)](https://godoc.org/github.com/izumin5210/newrelic-contrib-go)
[![Go project version](https://badge.fury.io/go/github.com%2Fizumin5210%2Fnewrelic-contrib-go.svg)](https://badge.fury.io/go/github.com%2Fizumin5210%2Fnewrelic-contrib-go)
[![license](https://img.shields.io/github/license/izumin5210/newrelic-contrib-go.svg)](./LICENSE)

Collection of utils and helpers to use [New Relic Agent](https://github.com/newrelic/go-agent) in production.

## Packages

- [`nrutil`](https://godoc.org/github.com/izumin5210/newrelic-contrib-go/nrutil)
  - a utilities to store [`newrelic.Transaction`](https://godoc.org/github.com/newrelic/go-agent#Transaction) into `context.Context`.
- [`nrhttp`](https://godoc.org/github.com/izumin5210/newrelic-contrib-go/nrhttp)
  - `http.RoundTripper` constructor to create an external segment of HTTP requests.
  - `http.Handler` wrappers to create a request segment.
- [`nrsql`](https://godoc.org/github.com/izumin5210/newrelic-contrib-go/nrsql)
  - Wrapper of `*sql.DB` to create database segments of queries.
