package flight

import (
	"sync"
)

// RefundPassengerFeeParam 结构体
type RefundPassengerFeeParam struct {
	// 乘机人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 机票已使用部分总价(单位:分)
	AlreadyUsedTotalPirce int64 `json:"already_used_total_pirce,omitempty" xml:"already_used_total_pirce,omitempty"`
	// 机票不可退税费(单位:分)
	NonRefundableTaxPrice int64 `json:"non_refundable_tax_price,omitempty" xml:"non_refundable_tax_price,omitempty"`
	// 机票不可退票价(单位:分)
	NonRefundableTicketPrice int64 `json:"non_refundable_ticket_price,omitempty" xml:"non_refundable_ticket_price,omitempty"`
	// 机票不可退改签服务费(单位:分)
	NonRefundableTotalChangeServiceFee int64 `json:"non_refundable_total_change_service_fee,omitempty" xml:"non_refundable_total_change_service_fee,omitempty"`
	// 机票不可退改签升舱费(单位:分)
	NonRefundableTotalChangeUpgradeFee int64 `json:"non_refundable_total_change_upgrade_fee,omitempty" xml:"non_refundable_total_change_upgrade_fee,omitempty"`
}

var poolRefundPassengerFeeParam = sync.Pool{
	New: func() any {
		return new(RefundPassengerFeeParam)
	},
}

// GetRefundPassengerFeeParam() 从对象池中获取RefundPassengerFeeParam
func GetRefundPassengerFeeParam() *RefundPassengerFeeParam {
	return poolRefundPassengerFeeParam.Get().(*RefundPassengerFeeParam)
}

// ReleaseRefundPassengerFeeParam 释放RefundPassengerFeeParam
func ReleaseRefundPassengerFeeParam(v *RefundPassengerFeeParam) {
	v.PassengerName = ""
	v.AlreadyUsedTotalPirce = 0
	v.NonRefundableTaxPrice = 0
	v.NonRefundableTicketPrice = 0
	v.NonRefundableTotalChangeServiceFee = 0
	v.NonRefundableTotalChangeUpgradeFee = 0
	poolRefundPassengerFeeParam.Put(v)
}
