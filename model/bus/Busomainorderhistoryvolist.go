package bus

// Busomainorderhistoryvolist 结构体
type Busomainorderhistoryvolist struct {
	// lastPlaceName 目的地
	LastPlaceName string `json:"last_place_name,omitempty" xml:"last_place_name,omitempty"`
	// scheduleId 车次id
	ScheduleId string `json:"schedule_id,omitempty" xml:"schedule_id,omitempty"`
	// gmtCreate 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// startCityName 出发城市
	StartCityName string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty"`
	// issueTime 出票时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// payTime 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// agentOrderId 代理商订单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// startTime 出发时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// endCityName 到达城市
	EndCityName string `json:"end_city_name,omitempty" xml:"end_city_name,omitempty"`
	// arriveStationName 到达站名称
	ArriveStationName string `json:"arrive_station_name,omitempty" xml:"arrive_station_name,omitempty"`
	// busNumber 车次编号
	BusNumber string `json:"bus_number,omitempty" xml:"bus_number,omitempty"`
	// startStationName 出发车站
	StartStationName string `json:"start_station_name,omitempty" xml:"start_station_name,omitempty"`
	// extAttr json格式的扩展信息，如自助机支付方式等
	ExtAttr string `json:"ext_attr,omitempty" xml:"ext_attr,omitempty"`
	// alipayTradeId 支付宝交易号
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
	// refundAmount 退款金额（分）
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// agentId 代理商编号
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// canRefund 是否可退 1 可退 0 不可退
	CanRefund int64 `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
	// mainOrderId 飞猪订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// orderStatus 订单状态. 10：初始状态，15:订单可见；41:订单完成待打款给卖家；70：交易成功 80:交易关闭
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// discountAmount 折扣优惠价格
	DiscountAmount int64 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// payStatus 支付状态 	支付状态(10:初始支付状态,20:创建担保交易成功,30:买家付款成功,60:已打款给卖家,100:关闭已支付订单成功,101:关闭未支付订单成功)
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// ticketCount 票数
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// totalPrice 总价
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// refundStatus 退款状态:0 无申请 10-初始 20-无需退款 40-退款成功
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// orderModel 1 自助机 0 线上订单
	OrderModel int64 `json:"order_model,omitempty" xml:"order_model,omitempty"`
}
