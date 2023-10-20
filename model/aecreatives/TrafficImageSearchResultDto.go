package aecreatives

import (
	"sync"
)

// TrafficImageSearchResultDto 结构体
type TrafficImageSearchResultDto struct {
	// 图搜结果
	Products []Product `json:"products,omitempty" xml:"products>product,omitempty"`
	// 总数
	TotalRecordCount int64 `json:"total_record_count,omitempty" xml:"total_record_count,omitempty"`
	// 图片识别的坐标
	Region *ProductImgRegionDto `json:"region,omitempty" xml:"region,omitempty"`
}

var poolTrafficImageSearchResultDto = sync.Pool{
	New: func() any {
		return new(TrafficImageSearchResultDto)
	},
}

// GetTrafficImageSearchResultDto() 从对象池中获取TrafficImageSearchResultDto
func GetTrafficImageSearchResultDto() *TrafficImageSearchResultDto {
	return poolTrafficImageSearchResultDto.Get().(*TrafficImageSearchResultDto)
}

// ReleaseTrafficImageSearchResultDto 释放TrafficImageSearchResultDto
func ReleaseTrafficImageSearchResultDto(v *TrafficImageSearchResultDto) {
	v.Products = v.Products[:0]
	v.TotalRecordCount = 0
	v.Region = nil
	poolTrafficImageSearchResultDto.Put(v)
}
