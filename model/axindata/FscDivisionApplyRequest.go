package axindata

import (
	"sync"
)

// FscDivisionApplyRequest 结构体
type FscDivisionApplyRequest struct {
	// 城市新增列表
	DivisionList []FscDivisionApplyApiDto `json:"division_list,omitempty" xml:"division_list>fsc_division_apply_api_dto,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolFscDivisionApplyRequest = sync.Pool{
	New: func() any {
		return new(FscDivisionApplyRequest)
	},
}

// GetFscDivisionApplyRequest() 从对象池中获取FscDivisionApplyRequest
func GetFscDivisionApplyRequest() *FscDivisionApplyRequest {
	return poolFscDivisionApplyRequest.Get().(*FscDivisionApplyRequest)
}

// ReleaseFscDivisionApplyRequest 释放FscDivisionApplyRequest
func ReleaseFscDivisionApplyRequest(v *FscDivisionApplyRequest) {
	v.DivisionList = v.DivisionList[:0]
	v.SupplierId = ""
	poolFscDivisionApplyRequest.Put(v)
}
