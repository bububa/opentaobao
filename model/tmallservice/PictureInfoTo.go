package tmallservice

// PictureInfoTo 
type PictureInfoTo struct {
    // pixel
    Pixel   string `json:"pixel,omitempty" xml:"pixel,omitempty"`
    // sizes
    Sizes   int64 `json:"sizes,omitempty" xml:"sizes,omitempty"`
    // pictureUrl
    PictureUrl   string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
}
