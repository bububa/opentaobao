package waybill

import (
	"sync"
)

// WaybillApplyRequest 结构体
type WaybillApplyRequest struct {
	// 物流服务商ID
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}

var poolWaybillApplyRequest = sync.Pool{
	New: func() any {
		return new(WaybillApplyRequest)
	},
}

// GetWaybillApplyRequest() 从对象池中获取WaybillApplyRequest
func GetWaybillApplyRequest() *WaybillApplyRequest {
	return poolWaybillApplyRequest.Get().(*WaybillApplyRequest)
}

// ReleaseWaybillApplyRequest 释放WaybillApplyRequest
func ReleaseWaybillApplyRequest(v *WaybillApplyRequest) {
	v.CpCode = ""
	poolWaybillApplyRequest.Put(v)
}
