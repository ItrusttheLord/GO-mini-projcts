package middlware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rw := &ResponseWriter{ResponseWriter: w}
	l.handler.ServeHTTP(rw, r)
	logrus.Infof("%s %s %d %v", r.Method, r.URL.Path, rw.StatusCode, time.Since(start))
}

func NewLogger(handlerToWrapp http.Handler) *Logger {
	return &Logger{handlerToWrapp}
}

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}
