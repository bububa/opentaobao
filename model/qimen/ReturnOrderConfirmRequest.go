package qimen

import (
	"sync"
)

// ReturnOrderConfirmRequest 结构体
type ReturnOrderConfirmRequest struct {
	// 订单信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 退货单信息
	ReturnOrder *ReturnOrder `json:"returnOrder,omitempty" xml:"returnOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenReturnorderConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolReturnOrderConfirmRequest = sync.Pool{
	New: func() any {
		return new(ReturnOrderConfirmRequest)
	},
}

// GetReturnOrderConfirmRequest() 从对象池中获取ReturnOrderConfirmRequest
func GetReturnOrderConfirmRequest() *ReturnOrderConfirmRequest {
	return poolReturnOrderConfirmRequest.Get().(*ReturnOrderConfirmRequest)
}

// ReleaseReturnOrderConfirmRequest 释放ReturnOrderConfirmRequest
func ReleaseReturnOrderConfirmRequest(v *ReturnOrderConfirmRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.ReturnOrder = nil
	v.ExtendProps = nil
	poolReturnOrderConfirmRequest.Put(v)
}
