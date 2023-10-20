package wenyuvideo

import (
	"sync"
)

// YoukuWenyuvideoSeetaGetModel 结构体
type YoukuWenyuvideoSeetaGetModel struct {
	// logo数组
	Logos []string `json:"logos,omitempty" xml:"logos>string,omitempty"`
	// 名字数组
	Names []string `json:"names,omitempty" xml:"names>string,omitempty"`
	// 片段数组
	Segments []Segments `json:"segments,omitempty" xml:"segments>segments,omitempty"`
}

var poolYoukuWenyuvideoSeetaGetModel = sync.Pool{
	New: func() any {
		return new(YoukuWenyuvideoSeetaGetModel)
	},
}

// GetYoukuWenyuvideoSeetaGetModel() 从对象池中获取YoukuWenyuvideoSeetaGetModel
func GetYoukuWenyuvideoSeetaGetModel() *YoukuWenyuvideoSeetaGetModel {
	return poolYoukuWenyuvideoSeetaGetModel.Get().(*YoukuWenyuvideoSeetaGetModel)
}

// ReleaseYoukuWenyuvideoSeetaGetModel 释放YoukuWenyuvideoSeetaGetModel
func ReleaseYoukuWenyuvideoSeetaGetModel(v *YoukuWenyuvideoSeetaGetModel) {
	v.Logos = v.Logos[:0]
	v.Names = v.Names[:0]
	v.Segments = v.Segments[:0]
	poolYoukuWenyuvideoSeetaGetModel.Put(v)
}
