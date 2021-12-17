var videoElement;
var adDisplayContainer;
var adsLoaded = false;
var adContainer;
var adsLoader;
var adsManager;
var userID = Math.floor(Math.random() * 10000); // ユーザID（今はランダム）

window.addEventListener("load", function () {
  videoElement = document.getElementById("video-element");

  // IMASDKの初期化処理
  initializeIMA();

  var playButton = document.getElementById("play-button");
  playButton.addEventListener("click", function (event) {
    videoElement.play();
  });

  videoElement.addEventListener("play", function (event) {
    if (adsLoaded) {
      return;
    }
    adsLoaded = true;

    event.preventDefault();

    videoElement.load();
    adDisplayContainer.initialize();

    // adManagerをスタートして広告を再生する
    try {
      adsManager.init(
        videoElement.clientWidth,
        videoElement.clientHeight,
        google.ima.ViewMode.NORMAL
      );
      adsManager.start();
    } catch (adError) {
      videoElement.play();
    }
  });
});

function initializeIMA() {
  // 動画画面をオーバーレイして広告を表示するためのコンテナ
  adContainer = document.getElementById("ad-container");

  // クリックで一時再生・停止
  // adContainerがvideo要素を覆ってしまうため、adContainerのクリックで一時停止・再生する
  adContainer.addEventListener("click", function (event) {
    if (videoElement.paused) {
      videoElement.play();
    } else {
      videoElement.pause();
    }
  });

  // adsLoaderを作成
  // 広告を取得して表示するオブジェクト
  adDisplayContainer = new google.ima.AdDisplayContainer(
    adContainer,
    videoElement
  );
  adsLoader = new google.ima.AdsLoader(adDisplayContainer);

  // 広告ロードしたらADS_MANAGER_LOADED eventからAdManagerを作成
  adsLoader.addEventListener(
    google.ima.AdsManagerLoadedEvent.Type.ADS_MANAGER_LOADED,
    function (event) {
      adsManager = event.getAdsManager(videoElement);

      // エラーハンドリング
      adsManager.addEventListener(
        google.ima.AdErrorEvent.Type.AD_ERROR,
        onAdError
      );

      // CONTENT_PAUSE_REQUESTEDで再生を一時停止
      // 裏にいる動画プレイヤーから発火されるevent
      adsManager.addEventListener(
        google.ima.AdEvent.Type.CONTENT_PAUSE_REQUESTED,
        function (event) {
          videoElement.pause();
        }
      );

      // CONTENT_PAUSE_REQUESTEDで再生を一時停止
      // 裏にいる動画プレイヤーから発火されるevent
      adsManager.addEventListener(
        google.ima.AdEvent.Type.CONTENT_RESUME_REQUESTED,
        function (event) {
          videoElement.play();
        }
      );
      adsManager.addEventListener(
        google.ima.AdEvent.Type.LOADED,
        function (event) {
          var ad = event.getAd();
          if (!ad.isLinear()) {
            videoElement.play();
          }
        }
      );
    },
    false
  );

  // ロードエラー時のハンドリング
  adsLoader.addEventListener(
    google.ima.AdErrorEvent.Type.AD_ERROR,
    onAdError,
    false
  );

  // Let the AdsLoader know when the video has ended
  videoElement.addEventListener("ended", function () {
    adsLoader.contentComplete();
  });

  var adsRequest = new google.ima.AdsRequest();
  adsRequest.adTagUrl = "http://localhost:8080/ad/vast?userID=" + userID;
  adsRequest.linearAdSlotWidth = videoElement.clientWidth;
  adsRequest.linearAdSlotHeight = videoElement.clientHeight;

  // 広告を取得する
  adsLoader.requestAds(adsRequest);
}

function onAdError(adErrorEvent) {
  console.error("Ad Error: " + adErrorEvent.getError());
  if (adsManager) {
    adsManager.destroy();
  }
}

// 動画コンテナのresizeに追従する
window.addEventListener("resize", function (event) {
  adsManager.resize(
    videoElement.clientWidth,
    videoElement.clientHeight,
    google.ima.ViewMode.NORMAL
  );
});
