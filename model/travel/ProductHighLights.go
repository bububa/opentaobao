package travel

import (
	"sync"
)

// ProductHighLights 结构体
type ProductHighLights struct {
	// 产品亮点图片
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
	// 产品亮点标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 产品亮点描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
}

var poolProductHighLights = sync.Pool{
	New: func() any {
		return new(ProductHighLights)
	},
}

// GetProductHighLights() 从对象池中获取ProductHighLights
func GetProductHighLights() *ProductHighLights {
	return poolProductHighLights.Get().(*ProductHighLights)
}

// ReleaseProductHighLights 释放ProductHighLights
func ReleaseProductHighLights(v *ProductHighLights) {
	v.PicUrls = v.PicUrls[:0]
	v.Title = ""
	v.Desc = ""
	poolProductHighLights.Put(v)
}
