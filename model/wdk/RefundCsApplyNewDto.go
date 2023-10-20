package wdk

import (
	"sync"
)

// RefundCsApplyNewDto 结构体
type RefundCsApplyNewDto struct {
	// 申请退款的子订单ID列表
	SubRefundOrders []CsApplySubOrderDto `json:"sub_refund_orders,omitempty" xml:"sub_refund_orders>cs_apply_sub_order_dto,omitempty"`
	// 渠道订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 商家经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 请求唯一键
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 备注说明
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 申请退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 退款原因id
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
	// 渠道来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 申请退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 申请退运费
	RefundDeliveryFee int64 `json:"refund_delivery_fee,omitempty" xml:"refund_delivery_fee,omitempty"`
	// 申请退包装费
	RefundPackageFee int64 `json:"refund_package_fee,omitempty" xml:"refund_package_fee,omitempty"`
}

var poolRefundCsApplyNewDto = sync.Pool{
	New: func() any {
		return new(RefundCsApplyNewDto)
	},
}

// GetRefundCsApplyNewDto() 从对象池中获取RefundCsApplyNewDto
func GetRefundCsApplyNewDto() *RefundCsApplyNewDto {
	return poolRefundCsApplyNewDto.Get().(*RefundCsApplyNewDto)
}

// ReleaseRefundCsApplyNewDto 释放RefundCsApplyNewDto
func ReleaseRefundCsApplyNewDto(v *RefundCsApplyNewDto) {
	v.SubRefundOrders = v.SubRefundOrders[:0]
	v.OutOrderId = ""
	v.StoreId = ""
	v.RequestId = ""
	v.Memo = ""
	v.RefundReason = ""
	v.ReasonId = 0
	v.OrderFrom = 0
	v.RefundFee = 0
	v.RefundDeliveryFee = 0
	v.RefundPackageFee = 0
	poolRefundCsApplyNewDto.Put(v)
}
