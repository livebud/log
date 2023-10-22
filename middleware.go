package log

import (
	"context"
	"fmt"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/segmentio/ksuid"
)

// ErrNotInContext is returned when a log is not in the context
var ErrNotInContext = fmt.Errorf("log: not in context")

type contextKey string

const logKey contextKey = "log"

// FromContext gets the log from the context
func FromContext(ctx context.Context) (Log, error) {
	log, ok := ctx.Value(logKey).(Log)
	if !ok {
		return nil, ErrNotInContext
	}
	return log, nil
}

// Middleware logging for an HTTP handler
func (l *Logger) Middleware() *Middleware {
	return &Middleware{
		Log: l,
		// RequestId is a function for generating a unique request id
		RequestId: func() string {
			return ksuid.New().String()
		},
	}
}

type Middleware struct {
	Log       Log
	RequestId func() string
}

func (m *Middleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Support an existing request id
		requestId := r.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = m.RequestId()
			r.Header.Set("X-Request-Id", requestId)
		}
		log := m.Log.Fields(Fields{
			"url":         r.RequestURI,
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			"request_id":  requestId,
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
}
