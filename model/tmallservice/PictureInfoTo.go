package tmallservice

// PictureInfoTo 结构体
type PictureInfoTo struct {
	// pictureUrl
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// pixel
	Pixel string `json:"pixel,omitempty" xml:"pixel,omitempty"`
	// sizes
	Sizes int64 `json:"sizes,omitempty" xml:"sizes,omitempty"`
}
