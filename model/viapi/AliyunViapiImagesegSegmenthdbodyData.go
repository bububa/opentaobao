package viapi

import (
	"sync"
)

// AliyunViapiImagesegSegmenthdbodyData 结构体
type AliyunViapiImagesegSegmenthdbodyData struct {
	// 输出图像URL
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

var poolAliyunViapiImagesegSegmenthdbodyData = sync.Pool{
	New: func() any {
		return new(AliyunViapiImagesegSegmenthdbodyData)
	},
}

// GetAliyunViapiImagesegSegmenthdbodyData() 从对象池中获取AliyunViapiImagesegSegmenthdbodyData
func GetAliyunViapiImagesegSegmenthdbodyData() *AliyunViapiImagesegSegmenthdbodyData {
	return poolAliyunViapiImagesegSegmenthdbodyData.Get().(*AliyunViapiImagesegSegmenthdbodyData)
}

// ReleaseAliyunViapiImagesegSegmenthdbodyData 释放AliyunViapiImagesegSegmenthdbodyData
func ReleaseAliyunViapiImagesegSegmenthdbodyData(v *AliyunViapiImagesegSegmenthdbodyData) {
	v.ImageUrl = ""
	poolAliyunViapiImagesegSegmenthdbodyData.Put(v)
}
