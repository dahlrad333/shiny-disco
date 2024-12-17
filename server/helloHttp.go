package helper

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type LoggingMiddleware struct {
	Handler http.Handler
}

func (lm *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	log.Printf("Started %s %s", r.Method, r.URL.Path)

	lm.Handler.ServeHTTP(w, r)

	duration := time.Since(startTime)
	log.Printf("Completed %s in %v", r.URL.Path, duration)
}

func NewLoggingMiddleware(handler http.Handler) *LoggingMiddleware {
	return &LoggingMiddleware{Handler: handler}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}