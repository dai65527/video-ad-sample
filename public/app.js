var videoElement;

window.addEventListener("load", function () {
  videoElement = document.getElementById("video-element");
  var playButton = document.getElementById("play-button");
  playButton.addEventListener("click", function (event) {
    console.log("play button clicked");
    videoElement.play();
  });
});
