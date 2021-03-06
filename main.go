package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dai65527/vast-server-sample/adserver"
	"github.com/dai65527/vast-server-sample/logger"
	"github.com/dai65527/vast-server-sample/trackingserver"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("public")))      // 動画配信サーバ
	mux.Handle("/ad/vast", &adserver.AdServer{})              // 広告配信サーバ
	mux.Handle("/tracking", &trackingserver.TrackingServer{}) // トラッキングサーバ

	err := http.ListenAndServe(":8080", logger.HttpLogger(mux))
	log.Printf("server error: %v", err)
	os.Exit(1)
}
