package tmallgenie

// PlayUrl 结构体
type PlayUrl struct {
	// 可播放链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 节目音视频类型，目前支持audio及video两种类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 码率
	Bitrate int64 `json:"bitrate,omitempty" xml:"bitrate,omitempty"`
}
