package viapi

import (
	"sync"
)

// AliyunViapiImagesegSegmentCommonimageData 结构体
type AliyunViapiImagesegSegmentCommonimageData struct {
	// 抠图结果（png透明图）有效期半个小时
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

var poolAliyunViapiImagesegSegmentCommonimageData = sync.Pool{
	New: func() any {
		return new(AliyunViapiImagesegSegmentCommonimageData)
	},
}

// GetAliyunViapiImagesegSegmentCommonimageData() 从对象池中获取AliyunViapiImagesegSegmentCommonimageData
func GetAliyunViapiImagesegSegmentCommonimageData() *AliyunViapiImagesegSegmentCommonimageData {
	return poolAliyunViapiImagesegSegmentCommonimageData.Get().(*AliyunViapiImagesegSegmentCommonimageData)
}

// ReleaseAliyunViapiImagesegSegmentCommonimageData 释放AliyunViapiImagesegSegmentCommonimageData
func ReleaseAliyunViapiImagesegSegmentCommonimageData(v *AliyunViapiImagesegSegmentCommonimageData) {
	v.ImageUrl = ""
	poolAliyunViapiImagesegSegmentCommonimageData.Put(v)
}
