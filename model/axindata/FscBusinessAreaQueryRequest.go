package axindata

import (
	"sync"
)

// FscBusinessAreaQueryRequest 结构体
type FscBusinessAreaQueryRequest struct {
	// 区域类型 1:出境 2:国内
	BusinessAreaRange string `json:"business_area_range,omitempty" xml:"business_area_range,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}

var poolFscBusinessAreaQueryRequest = sync.Pool{
	New: func() any {
		return new(FscBusinessAreaQueryRequest)
	},
}

// GetFscBusinessAreaQueryRequest() 从对象池中获取FscBusinessAreaQueryRequest
func GetFscBusinessAreaQueryRequest() *FscBusinessAreaQueryRequest {
	return poolFscBusinessAreaQueryRequest.Get().(*FscBusinessAreaQueryRequest)
}

// ReleaseFscBusinessAreaQueryRequest 释放FscBusinessAreaQueryRequest
func ReleaseFscBusinessAreaQueryRequest(v *FscBusinessAreaQueryRequest) {
	v.BusinessAreaRange = ""
	v.SupplierId = ""
	v.PageSize = 0
	v.PageIndex = 0
	poolFscBusinessAreaQueryRequest.Put(v)
}
