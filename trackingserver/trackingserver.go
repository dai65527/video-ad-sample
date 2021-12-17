package trackingserver

import (
	"log"
	"net/http"
)

type TrackingServer struct{}

var _ http.Handler = (*TrackingServer)(nil)

// ServeHTTP 広告視聴ログを記録
func (s *TrackingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// クエリパラメタを取得
	adID := r.URL.Query().Get("adID")
	userID := r.URL.Query().Get("userID")
	event := r.URL.Query().Get("event")

	// 広告視聴ログを記録
	log.Println("[ad watch log]", "ad", adID, "user", userID, "event", event)
}
