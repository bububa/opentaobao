package tmallservice

import (
	"sync"
)

// PictureInfoTo 结构体
type PictureInfoTo struct {
	// pictureUrl
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// pixel
	Pixel string `json:"pixel,omitempty" xml:"pixel,omitempty"`
	// sizes
	Sizes int64 `json:"sizes,omitempty" xml:"sizes,omitempty"`
}

var poolPictureInfoTo = sync.Pool{
	New: func() any {
		return new(PictureInfoTo)
	},
}

// GetPictureInfoTo() 从对象池中获取PictureInfoTo
func GetPictureInfoTo() *PictureInfoTo {
	return poolPictureInfoTo.Get().(*PictureInfoTo)
}

// ReleasePictureInfoTo 释放PictureInfoTo
func ReleasePictureInfoTo(v *PictureInfoTo) {
	v.PictureUrl = ""
	v.Pixel = ""
	v.Sizes = 0
	poolPictureInfoTo.Put(v)
}
