package axindata

import (
	"sync"
)

// FscTripDivisionApiResponse 结构体
type FscTripDivisionApiResponse struct {
	// 返回数据
	Data []FscTripDivisionApiDto `json:"data,omitempty" xml:"data>fsc_trip_division_api_dto,omitempty"`
	// 数据总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}

var poolFscTripDivisionApiResponse = sync.Pool{
	New: func() any {
		return new(FscTripDivisionApiResponse)
	},
}

// GetFscTripDivisionApiResponse() 从对象池中获取FscTripDivisionApiResponse
func GetFscTripDivisionApiResponse() *FscTripDivisionApiResponse {
	return poolFscTripDivisionApiResponse.Get().(*FscTripDivisionApiResponse)
}

// ReleaseFscTripDivisionApiResponse 释放FscTripDivisionApiResponse
func ReleaseFscTripDivisionApiResponse(v *FscTripDivisionApiResponse) {
	v.Data = v.Data[:0]
	v.Total = 0
	v.PageSize = 0
	v.PageIndex = 0
	poolFscTripDivisionApiResponse.Put(v)
}
