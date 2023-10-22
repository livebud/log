package log_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/livebud/log"
	"github.com/matryer/is"
)

func TestMiddleware(t *testing.T) {
	is := is.New(t)
	buffer := log.Buffer()
	logger := log.New(buffer)
	middleware := logger.Middleware()
	handler := middleware.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log, err := log.FromContext(r.Context())
		is.NoErr(err)
		log.Field("cool", "story").Info("hello")
		w.Write([]byte("hello world"))
	}))
	req := httptest.NewRequest("GET", "http://livebud.com/docs", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	is.Equal(rec.Code, http.StatusOK)
	is.Equal(rec.Body.String(), "hello world")
	entries := buffer.Entries()
	is.Equal(len(entries), 3)
	is.Equal(entries[0].Message, "request")
	fields := entries[0].Fields
	is.Equal(len(fields), 4)
	is.Equal(fields.Get("method"), "GET")
	is.Equal(fields.Get("url"), "http://livebud.com/docs")
	is.Equal(fields.Get("remote_addr"), "192.0.2.1:1234")
	requestId := fields.Get("request_id").(string)
	is.Equal(len(requestId), 27)
	is.Equal(entries[1].Message, "hello")
	fields = entries[1].Fields
	is.Equal(len(fields), 5)
	is.Equal(fields.Get("method"), "GET")
	is.Equal(fields.Get("url"), "http://livebud.com/docs")
	is.Equal(fields.Get("remote_addr"), "192.0.2.1:1234")
	is.Equal(fields.Get("request_id"), requestId)
	is.Equal(fields.Get("cool"), "story")
	is.Equal(entries[2].Message, "response")
	fields = entries[2].Fields
	is.Equal(len(fields), 7)
	is.Equal(fields.Get("method"), "GET")
	is.Equal(fields.Get("url"), "http://livebud.com/docs")
	is.Equal(fields.Get("remote_addr"), "192.0.2.1:1234")
	is.Equal(fields.Get("request_id"), requestId)
	is.Equal(fields.Get("status"), 200)
	is.Equal(fields.Get("size"), int64(11))
	is.True(fields.Get("duration").(time.Duration) > 0)
	fmt.Println(fields.Get("duration"))
}
