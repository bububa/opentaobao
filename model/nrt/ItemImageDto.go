package nrt

// ItemImageDto 结构体
type ItemImageDto struct {
	// 图片地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片次序
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
}
