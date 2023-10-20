package iot

import (
	"sync"
)

// ImageUrlParam 结构体
type ImageUrlParam struct {
	// 默认图片
	Img string `json:"img,omitempty" xml:"img,omitempty"`
	// 大图
	Large string `json:"large,omitempty" xml:"large,omitempty"`
	// 中图
	Medium string `json:"medium,omitempty" xml:"medium,omitempty"`
	// 小图
	Small string `json:"small,omitempty" xml:"small,omitempty"`
}

var poolImageUrlParam = sync.Pool{
	New: func() any {
		return new(ImageUrlParam)
	},
}

// GetImageUrlParam() 从对象池中获取ImageUrlParam
func GetImageUrlParam() *ImageUrlParam {
	return poolImageUrlParam.Get().(*ImageUrlParam)
}

// ReleaseImageUrlParam 释放ImageUrlParam
func ReleaseImageUrlParam(v *ImageUrlParam) {
	v.Img = ""
	v.Large = ""
	v.Medium = ""
	v.Small = ""
	poolImageUrlParam.Put(v)
}
