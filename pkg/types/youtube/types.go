package youtubedr

type VideoFormat struct {
	Itag             int    `json:"iTag"`
	URL              string `json:"url"`
	MimeType         string `json:"mimeType"`
	Quality          string `json:"quality"`
	Cipher           string `json:"signatureCipher"`
	Bitrate          int    `json:"bitrate"`
	FPS              int    `json:"fps"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	LastModified     string `json:"lastModified"`
	ContentLength    int64  `json:"contentLength,string"`
	ProjectionType   string `json:"projectionType"`
	VideoQuality     string `json:"videoQuality"`
	Size             int64  `json:"size"`
	AverageBitrate   int    `json:"averageBitrate"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
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
