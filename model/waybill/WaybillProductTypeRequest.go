package waybill

import (
	"sync"
)

// WaybillProductTypeRequest 结构体
type WaybillProductTypeRequest struct {
	// 物流商编码CODE
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}

var poolWaybillProductTypeRequest = sync.Pool{
	New: func() any {
		return new(WaybillProductTypeRequest)
	},
}

// GetWaybillProductTypeRequest() 从对象池中获取WaybillProductTypeRequest
func GetWaybillProductTypeRequest() *WaybillProductTypeRequest {
	return poolWaybillProductTypeRequest.Get().(*WaybillProductTypeRequest)
}

// ReleaseWaybillProductTypeRequest 释放WaybillProductTypeRequest
func ReleaseWaybillProductTypeRequest(v *WaybillProductTypeRequest) {
	v.CpCode = ""
	poolWaybillProductTypeRequest.Put(v)
}
