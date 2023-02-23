package miniapp

// Image 结构体
type Image struct {
	// 图片链接
	IconUrl string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// 宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}
