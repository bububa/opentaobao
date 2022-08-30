package bus

// Merchantbusorderinfo 结构体
type Merchantbusorderinfo struct {
	// 商家订单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 商家侧支付宝流水号
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
	// 出发站
	StartStation string `json:"start_station,omitempty" xml:"start_station,omitempty"`
	// 到达站
	EndStation string `json:"end_station,omitempty" xml:"end_station,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 出票时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 交易类型
	TransType string `json:"trans_type,omitempty" xml:"trans_type,omitempty"`
	// 飞猪订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 订单状态，待支付：01；已取消：02；待出票：03；出票成功：04；已关闭：05
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 总金额，单位分
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 票数
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 车次信息
	BusLineInfo *MerchantBusLineInfo `json:"bus_line_info,omitempty" xml:"bus_line_info,omitempty"`
	// 是否是阿斯兰模式
	IsSelfSaleOrder bool `json:"is_self_sale_order,omitempty" xml:"is_self_sale_order,omitempty"`
	// 是否是微信渠道
	IsWechat bool `json:"is_wechat,omitempty" xml:"is_wechat,omitempty"`
}
