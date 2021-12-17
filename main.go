package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dai65527/vast-server-sample/adserver"
	"github.com/dai65527/vast-server-sample/logger"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.Handle("/ad/vast", &adserver.AdServer{})

	err := http.ListenAndServe(":8080", logger.HttpLogger(mux))
	log.Printf("server error: %v", err)
	os.Exit(1)
}
