package viapi

import (
	"sync"
)

// AliyunViapiObjectdetDetectobjectData 结构体
type AliyunViapiObjectdetDetectobjectData struct {
	// 人体检测框的集合
	Elements []Elements `json:"elements,omitempty" xml:"elements>elements,omitempty"`
	// 输入图片的高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 输入图片的宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}

var poolAliyunViapiObjectdetDetectobjectData = sync.Pool{
	New: func() any {
		return new(AliyunViapiObjectdetDetectobjectData)
	},
}

// GetAliyunViapiObjectdetDetectobjectData() 从对象池中获取AliyunViapiObjectdetDetectobjectData
func GetAliyunViapiObjectdetDetectobjectData() *AliyunViapiObjectdetDetectobjectData {
	return poolAliyunViapiObjectdetDetectobjectData.Get().(*AliyunViapiObjectdetDetectobjectData)
}

// ReleaseAliyunViapiObjectdetDetectobjectData 释放AliyunViapiObjectdetDetectobjectData
func ReleaseAliyunViapiObjectdetDetectobjectData(v *AliyunViapiObjectdetDetectobjectData) {
	v.Elements = v.Elements[:0]
	v.Height = 0
	v.Width = 0
	poolAliyunViapiObjectdetDetectobjectData.Put(v)
}
