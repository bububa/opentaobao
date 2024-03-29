package flight

import (
	"sync"
)

// RefundOrderFillConfirmFeeRq 结构体
type RefundOrderFillConfirmFeeRq struct {
	// 回填费用参数列表
	FeeParams []RefundPassengerFeeParam `json:"fee_params,omitempty" xml:"fee_params>refund_passenger_fee_param,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 申请单ID
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 商家退款金额(单位：分)，用于审计金额正确性
	TotalSellerRefundMoney int64 `json:"total_seller_refund_money,omitempty" xml:"total_seller_refund_money,omitempty"`
}

var poolRefundOrderFillConfirmFeeRq = sync.Pool{
	New: func() any {
		return new(RefundOrderFillConfirmFeeRq)
	},
}

// GetRefundOrderFillConfirmFeeRq() 从对象池中获取RefundOrderFillConfirmFeeRq
func GetRefundOrderFillConfirmFeeRq() *RefundOrderFillConfirmFeeRq {
	return poolRefundOrderFillConfirmFeeRq.Get().(*RefundOrderFillConfirmFeeRq)
}

// ReleaseRefundOrderFillConfirmFeeRq 释放RefundOrderFillConfirmFeeRq
func ReleaseRefundOrderFillConfirmFeeRq(v *RefundOrderFillConfirmFeeRq) {
	v.FeeParams = v.FeeParams[:0]
	v.AgentId = 0
	v.RefundOrderId = 0
	v.TotalSellerRefundMoney = 0
	poolRefundOrderFillConfirmFeeRq.Put(v)
}
