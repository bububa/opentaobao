package aedropshiper

import (
	"sync"
)

// TrafficImageSearchResultDto 结构体
type TrafficImageSearchResultDto struct {
	// products
	Products []TrafficImageProductDto `json:"products,omitempty" xml:"products>traffic_image_product_dto,omitempty"`
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
	poolTrafficImageSearchResultDto.Put(v)
}
