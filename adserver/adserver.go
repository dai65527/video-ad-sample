package adserver

import "net/http"

// AdServer is a vast server that implements http.Handler
type AdServer struct{}

var _ http.Handler = (*AdServer)(nil)

// ServerHTTP is the interface for the HTTP server
func (s *AdServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
