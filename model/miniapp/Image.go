package miniapp

import (
	"sync"
)

// Image 结构体
type Image struct {
	// 图片链接
	IconUrl string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// 宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}

var poolImage = sync.Pool{
	New: func() any {
		return new(Image)
	},
}

// GetImage() 从对象池中获取Image
func GetImage() *Image {
	return poolImage.Get().(*Image)
}

// ReleaseImage 释放Image
func ReleaseImage(v *Image) {
	v.IconUrl = ""
	v.Width = 0
	v.Height = 0
	poolImage.Put(v)
}
