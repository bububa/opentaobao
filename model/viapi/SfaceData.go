package viapi

import (
	"sync"
)

// SfaceData 结构体
type SfaceData struct {
	// 识别出的人脸信息，具体结构描述见face
	Faces []Face `json:"faces,omitempty" xml:"faces>face,omitempty"`
	// 人脸区域高度
	Heihght int64 `json:"heihght,omitempty" xml:"heihght,omitempty"`
	// 人脸区域宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 以图片左上角为坐标原点，人脸区域左上角到x轴距离
	Y int64 `json:"y,omitempty" xml:"y,omitempty"`
	// 以图片左上角为坐标原点，人脸区域左上角到y轴距离
	X int64 `json:"x,omitempty" xml:"x,omitempty"`
}

var poolSfaceData = sync.Pool{
	New: func() any {
		return new(SfaceData)
	},
}

// GetSfaceData() 从对象池中获取SfaceData
func GetSfaceData() *SfaceData {
	return poolSfaceData.Get().(*SfaceData)
}

// ReleaseSfaceData 释放SfaceData
func ReleaseSfaceData(v *SfaceData) {
	v.Faces = v.Faces[:0]
	v.Heihght = 0
	v.Width = 0
	v.Y = 0
	v.X = 0
	poolSfaceData.Put(v)
}
