package waybill

import (
	"sync"
)

// WaybillDetailQueryByWaybillCodeRequest 结构体
type WaybillDetailQueryByWaybillCodeRequest struct {
	// 快递公司code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 请求id
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
	// 电子面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
}

var poolWaybillDetailQueryByWaybillCodeRequest = sync.Pool{
	New: func() any {
		return new(WaybillDetailQueryByWaybillCodeRequest)
	},
}

// GetWaybillDetailQueryByWaybillCodeRequest() 从对象池中获取WaybillDetailQueryByWaybillCodeRequest
func GetWaybillDetailQueryByWaybillCodeRequest() *WaybillDetailQueryByWaybillCodeRequest {
	return poolWaybillDetailQueryByWaybillCodeRequest.Get().(*WaybillDetailQueryByWaybillCodeRequest)
}

// ReleaseWaybillDetailQueryByWaybillCodeRequest 释放WaybillDetailQueryByWaybillCodeRequest
func ReleaseWaybillDetailQueryByWaybillCodeRequest(v *WaybillDetailQueryByWaybillCodeRequest) {
	v.CpCode = ""
	v.ObjectId = ""
	v.WaybillCode = ""
	poolWaybillDetailQueryByWaybillCodeRequest.Put(v)
}
