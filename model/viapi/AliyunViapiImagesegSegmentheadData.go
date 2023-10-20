package viapi

import (
	"sync"
)

// AliyunViapiImagesegSegmentheadData 结构体
type AliyunViapiImagesegSegmentheadData struct {
	// 人体检测框的集合
	Elements []Elements `json:"elements,omitempty" xml:"elements>elements,omitempty"`
}

var poolAliyunViapiImagesegSegmentheadData = sync.Pool{
	New: func() any {
		return new(AliyunViapiImagesegSegmentheadData)
	},
}

// GetAliyunViapiImagesegSegmentheadData() 从对象池中获取AliyunViapiImagesegSegmentheadData
func GetAliyunViapiImagesegSegmentheadData() *AliyunViapiImagesegSegmentheadData {
	return poolAliyunViapiImagesegSegmentheadData.Get().(*AliyunViapiImagesegSegmentheadData)
}

// ReleaseAliyunViapiImagesegSegmentheadData 释放AliyunViapiImagesegSegmentheadData
func ReleaseAliyunViapiImagesegSegmentheadData(v *AliyunViapiImagesegSegmentheadData) {
	v.Elements = v.Elements[:0]
	poolAliyunViapiImagesegSegmentheadData.Put(v)
}
