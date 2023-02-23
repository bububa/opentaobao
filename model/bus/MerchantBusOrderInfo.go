package bus

// MerchantBusOrderInfo 结构体
type MerchantBusOrderInfo struct {
	// 票信息列表
	BusTicketInfoList []MerchantBusTicketInfo `json:"bus_ticket_info_list,omitempty" xml:"bus_ticket_info_list>merchant_bus_ticket_info,omitempty"`
	// 到达站点
	EndStation string `json:"end_station,omitempty" xml:"end_station,omitempty"`
	// 出票时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 代理商订单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 支付流水号
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 业务类型：普通巴士，旅游巴士
	TransType string `json:"trans_type,omitempty" xml:"trans_type,omitempty"`
	// 出发站点
	StartStation string `json:"start_station,omitempty" xml:"start_station,omitempty"`
	// 票数量
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 票总价，单位分
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 订单状态：1-待支付；2-已取消；3-待出票；4-出票成功；5-已关闭
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 班次信息
	BusLineInfo *MerchantBusLineInfo `json:"bus_line_info,omitempty" xml:"bus_line_info,omitempty"`
	// 取票人信息
	BusFetchHolderInfo *MerchantBusFetchHolderInfo `json:"bus_fetch_holder_info,omitempty" xml:"bus_fetch_holder_info,omitempty"`
	// 飞猪订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 淘宝订单号
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 自营订单
	IsSelfSaleOrder bool `json:"is_self_sale_order,omitempty" xml:"is_self_sale_order,omitempty"`
	// 微信订单
	IsWechat bool `json:"is_wechat,omitempty" xml:"is_wechat,omitempty"`
}
