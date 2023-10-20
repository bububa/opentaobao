package waybill

import (
	"sync"
)

// WaybillDetailQueryByBizSubCodeRequest 结构体
type WaybillDetailQueryByBizSubCodeRequest struct {
	// 订单号
	BizSubCode string `json:"biz_sub_code,omitempty" xml:"biz_sub_code,omitempty"`
	// 请求id
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
}

var poolWaybillDetailQueryByBizSubCodeRequest = sync.Pool{
	New: func() any {
		return new(WaybillDetailQueryByBizSubCodeRequest)
	},
}

// GetWaybillDetailQueryByBizSubCodeRequest() 从对象池中获取WaybillDetailQueryByBizSubCodeRequest
func GetWaybillDetailQueryByBizSubCodeRequest() *WaybillDetailQueryByBizSubCodeRequest {
	return poolWaybillDetailQueryByBizSubCodeRequest.Get().(*WaybillDetailQueryByBizSubCodeRequest)
}

// ReleaseWaybillDetailQueryByBizSubCodeRequest 释放WaybillDetailQueryByBizSubCodeRequest
func ReleaseWaybillDetailQueryByBizSubCodeRequest(v *WaybillDetailQueryByBizSubCodeRequest) {
	v.BizSubCode = ""
	v.ObjectId = ""
	poolWaybillDetailQueryByBizSubCodeRequest.Put(v)
}
