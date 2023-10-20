package axindata

import (
	"sync"
)

// FscBusinessAreaApiDto 结构体
type FscBusinessAreaApiDto struct {
	// 业务区域ID
	BusinessAreaId string `json:"business_area_id,omitempty" xml:"business_area_id,omitempty"`
	// 业务区域名称
	BusinessAreaName string `json:"business_area_name,omitempty" xml:"business_area_name,omitempty"`
	// 区域类型 1:出境 2:国内
	BusinessAreaRange string `json:"business_area_range,omitempty" xml:"business_area_range,omitempty"`
}

var poolFscBusinessAreaApiDto = sync.Pool{
	New: func() any {
		return new(FscBusinessAreaApiDto)
	},
}

// GetFscBusinessAreaApiDto() 从对象池中获取FscBusinessAreaApiDto
func GetFscBusinessAreaApiDto() *FscBusinessAreaApiDto {
	return poolFscBusinessAreaApiDto.Get().(*FscBusinessAreaApiDto)
}

// ReleaseFscBusinessAreaApiDto 释放FscBusinessAreaApiDto
func ReleaseFscBusinessAreaApiDto(v *FscBusinessAreaApiDto) {
	v.BusinessAreaId = ""
	v.BusinessAreaName = ""
	v.BusinessAreaRange = ""
	poolFscBusinessAreaApiDto.Put(v)
}
