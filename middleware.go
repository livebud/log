package log

import (
	"context"
	"fmt"
	"net/http"

	"github.com/felixge/httpsnoop"
)

// ErrNotInContext is returned when a log is not in the context
var ErrNotInContext = fmt.Errorf("log: not in context")

type contextKey string

const logKey contextKey = "log"

// From gets the log from the context. If the logger isn't in the middleware,
// we warn and discards the logs
func From(ctx context.Context) Log {
	log, ok := ctx.Value(logKey).(Log)
	if !ok {
		Warn("log: not in context, discarding logs")
		return Discard()
	}
	return log
}

func (l *Logger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Support an existing request id
		requestId := r.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = l.requestId()
			r.Header.Set("X-Request-Id", requestId)
		}
		log := l.Fields(Fields{
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
