package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Custom ResponseWriter to capture status code
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//wrap the responseWriter
	rw := &ResponseWriter{ResponseWriter: w}
	//call the next handler
	l.handler.ServeHTTP(rw, r)
	logrus.Infof("%s %s %d %v", r.Method, r.URL.Path, rw.statusCode, time.Since(start)) //log request details
}

// create a new logger middleware
func NewLogger(handlerToWrapp http.Handler) *Logger {
	return &Logger{handlerToWrapp}
}
