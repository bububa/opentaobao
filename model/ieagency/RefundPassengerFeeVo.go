package ieagency

import (
	"sync"
)

// RefundPassengerFeeVo 结构体
type RefundPassengerFeeVo struct {
	// 活动列表
	RefundActivityVos []RefundActivityVo `json:"refund_activity_vos,omitempty" xml:"refund_activity_vos>refund_activity_vo,omitempty"`
	// 已使用票总价(单位:分)
	AlreadyUsedTotalPirce int64 `json:"already_used_total_pirce,omitempty" xml:"already_used_total_pirce,omitempty"`
	// 改签不可退服务费(单位：分)
	NonRefundableChangeServiceFee int64 `json:"non_refundable_change_service_fee,omitempty" xml:"non_refundable_change_service_fee,omitempty"`
	// 改签不可退升舱费(单位:分)
	NonRefundableChangeUpgradeFee int64 `json:"non_refundable_change_upgrade_fee,omitempty" xml:"non_refundable_change_upgrade_fee,omitempty"`
	// 机票不可退税费(单位:分)
	NonRefundableTaxPrice int64 `json:"non_refundable_tax_price,omitempty" xml:"non_refundable_tax_price,omitempty"`
	// 机票不可退票价(单位:分)
	NonRefundableTicketPrice int64 `json:"non_refundable_ticket_price,omitempty" xml:"non_refundable_ticket_price,omitempty"`
	// 乘机人ID
	PassengerId int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// 乘机人退总金额(单位:分)
	RefundToBuyerMoney int64 `json:"refund_to_buyer_money,omitempty" xml:"refund_to_buyer_money,omitempty"`
	// 乘机人红包收回(单位:分)
	TakeBackActivityMoney int64 `json:"take_back_activity_money,omitempty" xml:"take_back_activity_money,omitempty"`
}

var poolRefundPassengerFeeVo = sync.Pool{
	New: func() any {
		return new(RefundPassengerFeeVo)
	},
}

// GetRefundPassengerFeeVo() 从对象池中获取RefundPassengerFeeVo
func GetRefundPassengerFeeVo() *RefundPassengerFeeVo {
	return poolRefundPassengerFeeVo.Get().(*RefundPassengerFeeVo)
}

// ReleaseRefundPassengerFeeVo 释放RefundPassengerFeeVo
func ReleaseRefundPassengerFeeVo(v *RefundPassengerFeeVo) {
	v.RefundActivityVos = v.RefundActivityVos[:0]
	v.AlreadyUsedTotalPirce = 0
	v.NonRefundableChangeServiceFee = 0
	v.NonRefundableChangeUpgradeFee = 0
	v.NonRefundableTaxPrice = 0
	v.NonRefundableTicketPrice = 0
	v.PassengerId = 0
	v.RefundToBuyerMoney = 0
	v.TakeBackActivityMoney = 0
	poolRefundPassengerFeeVo.Put(v)
}
