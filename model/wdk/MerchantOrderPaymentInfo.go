package wdk

// MerchantOrderPaymentInfo 结构体
type MerchantOrderPaymentInfo struct {
	// 支付类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 付款金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 线上支付订单号，线下流水号，代金券/优惠券为优惠券实例id
	SerialNum string `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
	// 支付宝支付方式后获得的userId
	Tuid string `json:"tuid,omitempty" xml:"tuid,omitempty"`
	// 其他支付方式支付后获得userId，例如：微信平台的openId
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
}
