package tmallservice

// PictureInfoTo 结构体
type PictureInfoTo struct {
	// pixel
	Pixel string `json:"pixel,omitempty" xml:"pixel,omitempty"`
	// pictureUrl
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// sizes
	Sizes int64 `json:"sizes,omitempty" xml:"sizes,omitempty"`
}
