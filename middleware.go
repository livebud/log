package log

import (
	"context"
	"fmt"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/livebud/middleware"
	"github.com/segmentio/ksuid"
)

// ErrNotInContext is returned when a log is not in the context
var ErrNotInContext = fmt.Errorf("log: not in context")

type contextKey string

const logKey contextKey = "log"

// From gets the log from the context. If the logger isn't in the middleware,
// we warn and discards the logs
func From(ctx context.Context) (Log, error) {
	log, ok := ctx.Value(logKey).(Log)
	if !ok {
		return nil, ErrNotInContext
	}
	return log, nil
}

// MustFrom gets the log from the context or panics
func MustFrom(ctx context.Context) Log {
	log, err := From(ctx)
	if err != nil {
		panic(err)
	}
	return log
}

// WithRequestId sets the request id function for generating a unique request id
// for each request
func WithRequestId(fn func(r *http.Request) string) func(*middlewareOption) {
	return func(opts *middlewareOption) {
		opts.requestId = fn
	}
}

type middlewareOption struct {
	requestId func(r *http.Request) string
}

// RequestId is a function for generating a unique request id
func defaultRequestId(r *http.Request) string {
	// Support an existing request id
	requestId := r.Header.Get("X-Request-Id")
	if requestId == "" {
		requestId = ksuid.New().String()
		// Set just in case we use it later
		r.Header.Set("X-Request-Id", requestId)
	}
	return requestId
}

// Middleware uses the logger to log requests and responses
func Middleware(log Log, options ...func(*middlewareOption)) middleware.Middleware {
	opts := &middlewareOption{
		requestId: defaultRequestId,
	}
	for _, option := range options {
		option(opts)
	}
	return middleware.Func(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := log.Fields(Fields{
				"url":         r.RequestURI,
				"method":      r.Method,
				"remote_addr": r.RemoteAddr,
				"request_id":  opts.requestId(r),
			})
			ctx := context.WithValue(r.Context(), logKey, log)
			r = r.WithContext(ctx)
			log.Info("request")
			res := httpsnoop.CaptureMetrics(next, w, r)
			log = log.Fields(Fields{
				"status":   res.Code,
				"duration": res.Duration.Milliseconds(),
				"size":     res.Written,
			})
			switch {
			case res.Code >= 500:
				log.Error("response")
			case res.Code >= 400:
				log.Warn("response")
			default:
				log.Info("response")
			}
		})
	})
}
