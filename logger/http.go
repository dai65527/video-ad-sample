package logger

import (
	"log"
	"net/http"
)

// HttpLogger リクエストのロガー
func HttpLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[http request]", "method", r.Method, "url", r.URL.String())
		h.ServeHTTP(w, r)
	})
}
