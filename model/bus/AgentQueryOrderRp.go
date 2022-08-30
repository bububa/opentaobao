package bus

// AgentQueryOrderRp 结构体
type AgentQueryOrderRp struct {
	// 订单详情集合
	MerchantBusOrderInfoList []Merchantbusorderinfo `json:"merchant_bus_order_info_list,omitempty" xml:"merchant_bus_order_info_list>merchantbusorderinfo,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
