package aliexpress

import (
	"sync"
)

// OffsitePostImageVo 结构体
type OffsitePostImageVo struct {
	// 图片链接
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 图片高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 图片宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}

var poolOffsitePostImageVo = sync.Pool{
	New: func() any {
		return new(OffsitePostImageVo)
	},
}

// GetOffsitePostImageVo() 从对象池中获取OffsitePostImageVo
func GetOffsitePostImageVo() *OffsitePostImageVo {
	return poolOffsitePostImageVo.Get().(*OffsitePostImageVo)
}

// ReleaseOffsitePostImageVo 释放OffsitePostImageVo
func ReleaseOffsitePostImageVo(v *OffsitePostImageVo) {
	v.ImageUrl = ""
	v.Height = 0
	v.Width = 0
	poolOffsitePostImageVo.Put(v)
}
