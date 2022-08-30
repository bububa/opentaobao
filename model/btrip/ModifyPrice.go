package btrip

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
