package icbu

import (
	"sync"
)

// MainImage 结构体
type MainImage struct {
	// alibaba图片中心的图片URL列表，请使用alibaba.icbu.photobank.upload接口上传图片
	Images []string `json:"images,omitempty" xml:"images>string,omitempty"`
	// 水印是否有边框，有边框(Y)或者无边框(N)
	WatermarkFrame string `json:"watermark_frame,omitempty" xml:"watermark_frame,omitempty"`
	// 水印位置，在中间(center)或者在底部(bottom)
	WatermarkPosition string `json:"watermark_position,omitempty" xml:"watermark_position,omitempty"`
	// 是否打水印，是(true)或否(false)
	Watermark bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

var poolMainImage = sync.Pool{
	New: func() any {
		return new(MainImage)
	},
}

// GetMainImage() 从对象池中获取MainImage
func GetMainImage() *MainImage {
	return poolMainImage.Get().(*MainImage)
}

// ReleaseMainImage 释放MainImage
func ReleaseMainImage(v *MainImage) {
	v.Images = v.Images[:0]
	v.WatermarkFrame = ""
	v.WatermarkPosition = ""
	v.Watermark = false
	poolMainImage.Put(v)
}
