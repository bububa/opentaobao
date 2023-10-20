package lsttrade

import (
	"sync"
)

// RefundSubOrderInfo 结构体
type RefundSubOrderInfo struct {
	// 退款数量
	RefundCount int64 `json:"refund_count,omitempty" xml:"refund_count,omitempty"`
	// 退款金额，单位分
	RefundPayment int64 `json:"refund_payment,omitempty" xml:"refund_payment,omitempty"`
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

var poolRefundSubOrderInfo = sync.Pool{
	New: func() any {
		return new(RefundSubOrderInfo)
	},
}

// GetRefundSubOrderInfo() 从对象池中获取RefundSubOrderInfo
func GetRefundSubOrderInfo() *RefundSubOrderInfo {
	return poolRefundSubOrderInfo.Get().(*RefundSubOrderInfo)
}

// ReleaseRefundSubOrderInfo 释放RefundSubOrderInfo
func ReleaseRefundSubOrderInfo(v *RefundSubOrderInfo) {
	v.RefundCount = 0
	v.RefundPayment = 0
	v.SubOrderId = 0
	poolRefundSubOrderInfo.Put(v)
}
