package waybill

import (
	"sync"
)

// WaybillOrderConfirmResponse 结构体
type WaybillOrderConfirmResponse struct {
	// 面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
}

var poolWaybillOrderConfirmResponse = sync.Pool{
	New: func() any {
		return new(WaybillOrderConfirmResponse)
	},
}

// GetWaybillOrderConfirmResponse() 从对象池中获取WaybillOrderConfirmResponse
func GetWaybillOrderConfirmResponse() *WaybillOrderConfirmResponse {
	return poolWaybillOrderConfirmResponse.Get().(*WaybillOrderConfirmResponse)
}

// ReleaseWaybillOrderConfirmResponse 释放WaybillOrderConfirmResponse
func ReleaseWaybillOrderConfirmResponse(v *WaybillOrderConfirmResponse) {
	v.WaybillCode = ""
	poolWaybillOrderConfirmResponse.Put(v)
}
