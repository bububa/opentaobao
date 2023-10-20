package viapi

import (
	"sync"
)

// AliyunViapiImagesegSegmentcomodityData 结构体
type AliyunViapiImagesegSegmentcomodityData struct {
	// 抠图结果（png透明图）有效期半个小时
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

var poolAliyunViapiImagesegSegmentcomodityData = sync.Pool{
	New: func() any {
		return new(AliyunViapiImagesegSegmentcomodityData)
	},
}

// GetAliyunViapiImagesegSegmentcomodityData() 从对象池中获取AliyunViapiImagesegSegmentcomodityData
func GetAliyunViapiImagesegSegmentcomodityData() *AliyunViapiImagesegSegmentcomodityData {
	return poolAliyunViapiImagesegSegmentcomodityData.Get().(*AliyunViapiImagesegSegmentcomodityData)
}

// ReleaseAliyunViapiImagesegSegmentcomodityData 释放AliyunViapiImagesegSegmentcomodityData
func ReleaseAliyunViapiImagesegSegmentcomodityData(v *AliyunViapiImagesegSegmentcomodityData) {
	v.ImageUrl = ""
	poolAliyunViapiImagesegSegmentcomodityData.Put(v)
}
