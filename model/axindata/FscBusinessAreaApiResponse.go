package axindata

import (
	"sync"
)

// FscBusinessAreaApiResponse 结构体
type FscBusinessAreaApiResponse struct {
	// 返回数据
	Data []FscBusinessAreaApiDto `json:"data,omitempty" xml:"data>fsc_business_area_api_dto,omitempty"`
	// 数据总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}

var poolFscBusinessAreaApiResponse = sync.Pool{
	New: func() any {
		return new(FscBusinessAreaApiResponse)
	},
}

// GetFscBusinessAreaApiResponse() 从对象池中获取FscBusinessAreaApiResponse
func GetFscBusinessAreaApiResponse() *FscBusinessAreaApiResponse {
	return poolFscBusinessAreaApiResponse.Get().(*FscBusinessAreaApiResponse)
}

// ReleaseFscBusinessAreaApiResponse 释放FscBusinessAreaApiResponse
func ReleaseFscBusinessAreaApiResponse(v *FscBusinessAreaApiResponse) {
	v.Data = v.Data[:0]
	v.Total = 0
	v.PageSize = 0
	v.PageIndex = 0
	poolFscBusinessAreaApiResponse.Put(v)
}
