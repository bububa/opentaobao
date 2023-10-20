package travel

import (
	"sync"
)

// HighLights 结构体
type HighLights struct {
	// 图片列表
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
	// 亮点标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 亮点描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
}

var poolHighLights = sync.Pool{
	New: func() any {
		return new(HighLights)
	},
}

// GetHighLights() 从对象池中获取HighLights
func GetHighLights() *HighLights {
	return poolHighLights.Get().(*HighLights)
}

// ReleaseHighLights 释放HighLights
func ReleaseHighLights(v *HighLights) {
	v.PicUrls = v.PicUrls[:0]
	v.Title = ""
	v.Desc = ""
	poolHighLights.Put(v)
}
