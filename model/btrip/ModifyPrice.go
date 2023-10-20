package btrip

import (
	"sync"
)

// ModifyPrice 结构体
type ModifyPrice struct {
	// 乘客类型
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 改签手续费
	UpgradeFee int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
	// 升舱差价
	UpgradePrice int64 `json:"upgrade_price,omitempty" xml:"upgrade_price,omitempty"`
	// 改签支付价格
	ChangePaymentPrice int64 `json:"change_payment_price,omitempty" xml:"change_payment_price,omitempty"`
	// 改签费
	ModifyPrice int64 `json:"modify_price,omitempty" xml:"modify_price,omitempty"`
	// 票价(分)
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

var poolModifyPrice = sync.Pool{
	New: func() any {
		return new(ModifyPrice)
	},
}

// GetModifyPrice() 从对象池中获取ModifyPrice
func GetModifyPrice() *ModifyPrice {
	return poolModifyPrice.Get().(*ModifyPrice)
}

// ReleaseModifyPrice 释放ModifyPrice
func ReleaseModifyPrice(v *ModifyPrice) {
	v.PassengerType = 0
	v.UpgradeFee = 0
	v.UpgradePrice = 0
	v.ChangePaymentPrice = 0
	v.ModifyPrice = 0
	v.TicketPrice = 0
	poolModifyPrice.Put(v)
}
