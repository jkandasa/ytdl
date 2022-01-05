package youtubedr

type VideoFormat struct {
	Itag          int    `json:"iTag"`
	FPS           int    `json:"fps"`
	VideoQuality  string `json:"videoQuality"`
	AudioQuality  string `json:"audioQuality"`
	AudioChannels int    `json:"audioChannels"`
	Size          int64  `json:"size"`
	Bitrate       int    `json:"bitrate"`
	MimeType      string `json:"mimeType"`
	URL           string `json:"url"`
}

type VideoInfo struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Author      string        `json:"author"`
	Duration    string        `json:"duration"`
	Description string        `json:"description"`
	Thumbnails  []Thumbnail   `json:"thumbnails"`
	Formats     []VideoFormat `json:"formats"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}
