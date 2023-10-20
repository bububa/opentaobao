package qimen

import (
	"sync"
)

// ReturnOrderCreateRequest 结构体
type ReturnOrderCreateRequest struct {
	// 订单信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 退货单信息
	ReturnOrder *ReturnOrder `json:"returnOrder,omitempty" xml:"returnOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenReturnorderCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolReturnOrderCreateRequest = sync.Pool{
	New: func() any {
		return new(ReturnOrderCreateRequest)
	},
}

// GetReturnOrderCreateRequest() 从对象池中获取ReturnOrderCreateRequest
func GetReturnOrderCreateRequest() *ReturnOrderCreateRequest {
	return poolReturnOrderCreateRequest.Get().(*ReturnOrderCreateRequest)
}

// ReleaseReturnOrderCreateRequest 释放ReturnOrderCreateRequest
func ReleaseReturnOrderCreateRequest(v *ReturnOrderCreateRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.ReturnOrder = nil
	v.ExtendProps = nil
	poolReturnOrderCreateRequest.Put(v)
}
