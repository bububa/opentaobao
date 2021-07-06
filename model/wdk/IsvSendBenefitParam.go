package wdk

// IsvSendBenefitParam 结构体
type IsvSendBenefitParam struct {
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部erp门店编码
	OutShopCode string `json:"out_shop_code,omitempty" xml:"out_shop_code,omitempty"`
	// 支付宝支付单号
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
	// 设备id，isv这边可能给到的是设备编号
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 设备类型。1是人工pos，2是自助pos
	DeviceType int64 `json:"device_type,omitempty" xml:"device_type,omitempty"`
}
