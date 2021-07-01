package servicecenter

// OrderConfirmQueryDto 结构体
type OrderConfirmQueryDto struct {
	// APPKEY，必填
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 内购服务的规格CODE，必填
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 周期单位(必选 数值1:年 2:月， 3：天)，必填
	CycUnit string `json:"cyc_unit,omitempty" xml:"cyc_unit,omitempty"`
	// 周期数量，必填
	CycNum string `json:"cyc_num,omitempty" xml:"cyc_num,omitempty"`
	// 使用该参数完成查询订单等操作，可选
	OutTradeCode string `json:"out_trade_code,omitempty" xml:"out_trade_code,omitempty"`
	// 设备类型，目前只支持PC，可选
	DeviceType string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	// 计量型服务的数量，如果是计量型内购服务，则必填
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
