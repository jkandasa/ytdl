package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	appUtils "github.com/jkandasa/ytdl/cmd/app/utils"
	ytTY "github.com/jkandasa/ytdl/pkg/types/youtube"
	"github.com/kkdai/youtube/v2"
)

// Youtube api
func RegisterYoutubeRoutes(router *mux.Router) {
	router.HandleFunc("/api/youtube/info", getVideoInfo).Methods(http.MethodGet)
}

func getVideoInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "url can not be empty", http.StatusBadRequest)
		return
	}

	client := youtube.Client{}
	video, err := client.GetVideo(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// copied from
	// https://github.com/kkdai/youtube/blob/master/cmd/youtubedr/info.go
	videoInfo := ytTY.VideoInfo{
		Title:       video.Title,
		Author:      video.Author,
		Duration:    video.Duration.String(),
		Description: video.Description,
		Thumbnails:  make([]ytTY.Thumbnail, 0),
	}

	// update Thumbnails
	for _, thumbnail := range video.Thumbnails {
		videoInfo.Thumbnails = append(videoInfo.Thumbnails, ytTY.Thumbnail{
			URL:    thumbnail.URL,
			Width:  thumbnail.Width,
			Height: thumbnail.Height,
		})
	}

	for _, format := range video.Formats {
		bitrate := format.AverageBitrate
		if bitrate == 0 {
			// Some formats don't have the average bitrate
			bitrate = format.Bitrate
		}

		size := format.ContentLength
		if size == 0 {
			// Some formats don't have this information
			size = int64(float64(bitrate) * video.Duration.Seconds() / 8)
		}

		videoInfo.Formats = append(videoInfo.Formats, ytTY.VideoFormat{
			Itag:          format.ItagNo,
			FPS:           format.FPS,
			VideoQuality:  format.QualityLabel,
			AudioQuality:  strings.ToLower(strings.TrimPrefix(format.AudioQuality, "AUDIO_QUALITY_")),
			AudioChannels: format.AudioChannels,
			Size:          size,
			Bitrate:       bitrate,
			MimeType:      format.MimeType,
			URL:           format.URL,
		})
	}

	appUtils.WriteJsonResponse(w, videoInfo)
}
