package openmall

import (
	"sync"
)

// TopItemImageVo 结构体
type TopItemImageVo struct {
	// 商品图片链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 商品属性，itemImages中不返回
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 图片放在第几张（多图时设置），propertyImages中不返回
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
	// 图片ID，itemImages中不返回
	ImageId int64 `json:"image_id,omitempty" xml:"image_id,omitempty"`
}

var poolTopItemImageVo = sync.Pool{
	New: func() any {
		return new(TopItemImageVo)
	},
}

// GetTopItemImageVo() 从对象池中获取TopItemImageVo
func GetTopItemImageVo() *TopItemImageVo {
	return poolTopItemImageVo.Get().(*TopItemImageVo)
}

// ReleaseTopItemImageVo 释放TopItemImageVo
func ReleaseTopItemImageVo(v *TopItemImageVo) {
	v.Url = ""
	v.Properties = ""
	v.Position = 0
	v.ImageId = 0
	poolTopItemImageVo.Put(v)
}
