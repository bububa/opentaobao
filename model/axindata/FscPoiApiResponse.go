package axindata

import (
	"sync"
)

// FscPoiApiResponse 结构体
type FscPoiApiResponse struct {
	// 返回数据
	Data []FscPoiApiDto `json:"data,omitempty" xml:"data>fsc_poi_api_dto,omitempty"`
	// 数据总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}

var poolFscPoiApiResponse = sync.Pool{
	New: func() any {
		return new(FscPoiApiResponse)
	},
}

// GetFscPoiApiResponse() 从对象池中获取FscPoiApiResponse
func GetFscPoiApiResponse() *FscPoiApiResponse {
	return poolFscPoiApiResponse.Get().(*FscPoiApiResponse)
}

// ReleaseFscPoiApiResponse 释放FscPoiApiResponse
func ReleaseFscPoiApiResponse(v *FscPoiApiResponse) {
	v.Data = v.Data[:0]
	v.Total = 0
	v.PageSize = 0
	v.PageIndex = 0
	poolFscPoiApiResponse.Put(v)
}
