package adserver

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/vast"
)

// AdServer 広告配信(VAST配信)サーバ
type AdServer struct{}

var _ http.Handler = (*AdServer)(nil)

const TrkBaseURL = "http://localhost:8080"
const CreativeBaseURL = "http://localhost:8080"

// ServerHTTP HTTPハンドラ
func (s *AdServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// cors対応
	w.Header().Set("Access-Control-Allow-Origin", "http://imasdk.googleapis.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// userIDの取得
	userID := r.URL.Query().Get("userID")

	// VASTを生成して返す
	w.Header().Set("Content-Type", "text/xml")
	w.Write(createVAST(userID))
}

// trackingURL 広告視聴ログ発火用のURLを生成
func trackingURL(adID, userID, event string) string {
	return fmt.Sprintf("%s/tracking?adID=%s&userID=%s&event=%s", TrkBaseURL, adID, userID, event)
}

// createVAST VASTのXMLを生成
func createVAST(userID string) []byte {
	// ユーザごとに最適な広告を選択（したい）
	adID := strconv.Itoa(rand.Int() % 10000)

	v := vast.VAST{
		Version: "3.0",
		Ads: []vast.Ad{
			{
				ID:       adID,
				Sequence: 1,

				// 広告表示に必要な情報
				InLine: &vast.InLine{
					// 配信サーバの名前
					AdSystem: &vast.AdSystem{
						Name: "sample vast server",
					},

					// 広告の名前
					AdTitle: vast.CDATAString{
						CDATA: "sample ad",
					},

					// 広告の表示ログ発火URL
					Impressions: []vast.Impression{
						{URI: trackingURL(adID, userID, "imp")},
					},

					// 広告クリエイティブの情報
					Creatives: []vast.Creative{
						{
							// Linearは動画内に挿入される動画広告
							Linear: &vast.Linear{
								// 動画広告の時間
								Duration: vast.Duration(time.Second * 10),

								// 動画広告の表示ログ発火URL
								// start/midpoint/completeでそれぞれ発火する
								TrackingEvents: []vast.Tracking{
									{Event: "start", URI: trackingURL(adID, userID, "start")},
									{Event: "midpoint", URI: trackingURL(adID, userID, "midpoint")},
									{Event: "complete", URI: trackingURL(adID, userID, "complete")},
								},

								// 動画広告クリエイティブののURL
								MediaFiles: []vast.MediaFile{
									{
										Delivery: "progressive",
										Type:     "video/mp4",
										URI:      CreativeBaseURL + "/ad/creative.mp4",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// XMLに変換
	b, _ := xml.MarshalIndent(v, "", "  ")
	return b
}
