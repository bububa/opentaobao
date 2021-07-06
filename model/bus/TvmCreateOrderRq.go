package bus

// TvmCreateOrderRq 结构体
type TvmCreateOrderRq struct {
	// 乘客
	Passengers []TvmPassengerVo `json:"passengers,omitempty" xml:"passengers>tvm_passenger_vo,omitempty"`
	// 分润账户明细列表，是个数组，有几个分账，写几个。 注意 只有需要分润到多账号才需要填，否则为空。分账总和等于订单总价。
	AccountInDetails []AccountInDetail `json:"account_in_details,omitempty" xml:"account_in_details>account_in_detail,omitempty"`
	// 代理商订单ID
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 终端机编号，每个商家终端编号要进行唯一标识
	MachineNumber string `json:"machine_number,omitempty" xml:"machine_number,omitempty"`
	// 渠道商支付宝ID
	ServiceProviderId string `json:"service_provider_id,omitempty" xml:"service_provider_id,omitempty"`
	// 交易场景来源 window(窗口) self(自助机）
	TradeSource string `json:"trade_source,omitempty" xml:"trade_source,omitempty"`
	// 创建订单截止时间，不传值默认使用出发时间字段 yyyy-mm-dd HH:mm:ss
	CreateDeadline string `json:"create_deadline,omitempty" xml:"create_deadline,omitempty"`
	// 取值范围 ALIPAY （飞猪渠道）; WECHAT（微信）; BANKCARD（银行卡）;CASH（现金）; OWN_ALIPAY（自身支付宝渠道，非飞猪渠道）；UNION（多码合一）
	PayMode string `json:"pay_mode,omitempty" xml:"pay_mode,omitempty"`
	// 票总数量
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 订单总价格，含服务费 (单位分)
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 班次信息
	TvmBusLineInfo *TvmBusLineInfo `json:"tvm_bus_line_info,omitempty" xml:"tvm_bus_line_info,omitempty"`
	// 出票超时时间(支付成功后开始计算) 单位:秒，如果商家不设置，平台默认超时时间为60天(自助机小程序订单使用)
	IssueTimeout int64 `json:"issue_timeout,omitempty" xml:"issue_timeout,omitempty"`
	// 支付超时时间(创建订单成功后开始计算) 单位:秒,到达指定时间后，平台进行关闭订单操作(自助机小程序订单使用)
	PayTimeout int64 `json:"pay_timeout,omitempty" xml:"pay_timeout,omitempty"`
	// true实名（传身份证）,必须传true后续进行程序强制校验
	RealName bool `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// true:切换为自助机小程序订单;false:普通自助机订单
	TvmOnline bool `json:"tvm_online,omitempty" xml:"tvm_online,omitempty"`
}
