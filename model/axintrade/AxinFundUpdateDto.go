package axintrade

// AxinFundUpdateDto 结构体
type AxinFundUpdateDto struct {
	// 请求版本号，用于幂等校验
	ReqVersion string `json:"req_version,omitempty" xml:"req_version,omitempty"`
	// 扩展属性
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 支付宝交易号
	AlipayOrderId string `json:"alipay_order_id,omitempty" xml:"alipay_order_id,omitempty"`
	// 资金单ID
	FundId int64 `json:"fund_id,omitempty" xml:"fund_id,omitempty"`
	// 订单ID
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
}
