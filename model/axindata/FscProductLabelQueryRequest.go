package axindata

import (
	"sync"
)

// FscProductLabelQueryRequest 结构体
type FscProductLabelQueryRequest struct {
	// 主题类目 INTL_GROUP_TRAVEL：出境线路 DOM_GROUP_TRAVEL：国内线路
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 父节点主题ID
	ParentLabelId string `json:"parent_label_id,omitempty" xml:"parent_label_id,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}

var poolFscProductLabelQueryRequest = sync.Pool{
	New: func() any {
		return new(FscProductLabelQueryRequest)
	},
}

// GetFscProductLabelQueryRequest() 从对象池中获取FscProductLabelQueryRequest
func GetFscProductLabelQueryRequest() *FscProductLabelQueryRequest {
	return poolFscProductLabelQueryRequest.Get().(*FscProductLabelQueryRequest)
}

// ReleaseFscProductLabelQueryRequest 释放FscProductLabelQueryRequest
func ReleaseFscProductLabelQueryRequest(v *FscProductLabelQueryRequest) {
	v.CategoryId = ""
	v.ParentLabelId = ""
	v.SupplierId = ""
	v.PageSize = 0
	v.PageIndex = 0
	poolFscProductLabelQueryRequest.Put(v)
}
