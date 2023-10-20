package tmallgenie

import (
	"sync"
)

// ImageUrl 结构体
type ImageUrl struct {
	// 中图链接(具体大小范围暂无标准，接入方自定)
	Mediam string `json:"mediam,omitempty" xml:"mediam,omitempty"`
	// 小图链接(具体大小范围暂无标准，接入方自定)
	Small string `json:"small,omitempty" xml:"small,omitempty"`
	// 大图链接(具体大小范围暂无标准，接入方自定)
	Large string `json:"large,omitempty" xml:"large,omitempty"`
	// 中图链接(具体大小范围暂无标准，接入方自定)
	Medium string `json:"medium,omitempty" xml:"medium,omitempty"`
	// 如果不知道放那个就传默认图片
	Img string `json:"img,omitempty" xml:"img,omitempty"`
}

var poolImageUrl = sync.Pool{
	New: func() any {
		return new(ImageUrl)
	},
}

// GetImageUrl() 从对象池中获取ImageUrl
func GetImageUrl() *ImageUrl {
	return poolImageUrl.Get().(*ImageUrl)
}

// ReleaseImageUrl 释放ImageUrl
func ReleaseImageUrl(v *ImageUrl) {
	v.Mediam = ""
	v.Small = ""
	v.Large = ""
	v.Medium = ""
	v.Img = ""
	poolImageUrl.Put(v)
}
