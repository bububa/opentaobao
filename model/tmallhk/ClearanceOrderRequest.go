package tmallhk

import (
	"sync"
)

// ClearanceOrderRequest 结构体
type ClearanceOrderRequest struct {
	// 交易主订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolClearanceOrderRequest = sync.Pool{
	New: func() any {
		return new(ClearanceOrderRequest)
	},
}

// GetClearanceOrderRequest() 从对象池中获取ClearanceOrderRequest
func GetClearanceOrderRequest() *ClearanceOrderRequest {
	return poolClearanceOrderRequest.Get().(*ClearanceOrderRequest)
}

// ReleaseClearanceOrderRequest 释放ClearanceOrderRequest
func ReleaseClearanceOrderRequest(v *ClearanceOrderRequest) {
	v.BizOrderId = 0
	poolClearanceOrderRequest.Put(v)
}
