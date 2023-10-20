package viapi

import (
	"sync"
)

// Face 结构体
type Face struct {
	// 人脸Id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 相似人物的名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 相似概率
	Rate int64 `json:"rate,omitempty" xml:"rate,omitempty"`
}

var poolFace = sync.Pool{
	New: func() any {
		return new(Face)
	},
}

// GetFace() 从对象池中获取Face
func GetFace() *Face {
	return poolFace.Get().(*Face)
}

// ReleaseFace 释放Face
func ReleaseFace(v *Face) {
	v.Id = ""
	v.Name = ""
	v.Rate = 0
	poolFace.Put(v)
}
