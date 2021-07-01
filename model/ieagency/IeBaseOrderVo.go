package ieagency

// IeBaseOrderVo 结构体
type IeBaseOrderVo struct {
	// 到达机场
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 全部航班号
	FlightNos string `json:"flight_nos,omitempty" xml:"flight_nos,omitempty"`
	// 交易订单id
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
	// 订单创建时间
	GmtCreateTime string `json:"gmt_create_time,omitempty" xml:"gmt_create_time,omitempty"`
	// 订单成功时间
	SuccessTime string `json:"success_time,omitempty" xml:"success_time,omitempty"`
	// 服务费
	TotalServicePrice int64 `json:"total_service_price,omitempty" xml:"total_service_price,omitempty"`
	// 回程出发日期
	InboundDepTime string `json:"inbound_dep_time,omitempty" xml:"inbound_dep_time,omitempty"`
	// 信用状态(Normal:正常,Frozen:已冻结,Unfrozen:已解冻)
	CreditStatus string `json:"credit_status,omitempty" xml:"credit_status,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 回程到达时间
	InboundArrTime string `json:"inbound_arr_time,omitempty" xml:"inbound_arr_time,omitempty"`
	// 转交易时间
	Pay2AgentTime string `json:"pay2_agent_time,omitempty" xml:"pay2_agent_time,omitempty"`
	// 预定hk时间
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 佣金状态(Init:初始状态,Failure:扣佣金失败,Success:扣佣金成功;)
	CommissionStatus string `json:"commission_status,omitempty" xml:"commission_status,omitempty"`
	// 总票价
	TotalTicketPrice int64 `json:"total_ticket_price,omitempty" xml:"total_ticket_price,omitempty"`
	// 活动金额
	TotalActivityRemitPrice int64 `json:"total_activity_remit_price,omitempty" xml:"total_activity_remit_price,omitempty"`
	// 总税费
	TotalTaxPrice int64 `json:"total_tax_price,omitempty" xml:"total_tax_price,omitempty"`
	// 出发时间
	OutboundDepTime string `json:"outbound_dep_time,omitempty" xml:"outbound_dep_time,omitempty"`
	// 订单总价(包含打包商品的价格)
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 红包价格
	TotalRedpackage int64 `json:"total_redpackage,omitempty" xml:"total_redpackage,omitempty"`
	// 出发机场
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 佣金金额
	CommissionPrice int64 `json:"commission_price,omitempty" xml:"commission_price,omitempty"`
	// 行程类型(OneWay:单程,RoundTrip:往返,MultiCity:多程)
	TripType string `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// 订单状态(Init:初始状态,BookSuccess:"预定成功,PaySuccess:付款成功,TicketSuccess:订单成功,Close:订单关闭)
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单关闭时间
	CloseTime string `json:"close_time,omitempty" xml:"close_time,omitempty"`
	// 支付状态(Init:初始状态;create_secure_pay:创建担保交易成功,PaySuccess:付款成功,TransferSuccess:转交易成功, PayedClosed:已支付订单关闭,NoPayClosed:未支付订单关闭 )
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 关闭类型(Init:正常(未关闭),BuyerClose:买家关闭,NoAVClose:舱位已售完,NoEnoughAV:剩余座位数不足,PayOvertime:付款超时,TicketOrverTime:出票超时,Other:其它)
	CloseType string `json:"close_type,omitempty" xml:"close_type,omitempty"`
	// 乘机人数量
	PassengerCount int64 `json:"passenger_count,omitempty" xml:"passenger_count,omitempty"`
	// 预计佣金金额
	EstCommissionPrice int64 `json:"est_commission_price,omitempty" xml:"est_commission_price,omitempty"`
	// 关闭原因
	CloseReason string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// 出发到达日期
	OutboundArrTime string `json:"outbound_arr_time,omitempty" xml:"outbound_arr_time,omitempty"`
	// 订单初始总价(包含打包商品的价格)
	OriginTotalPrice int64 `json:"origin_total_price,omitempty" xml:"origin_total_price,omitempty"`
	// 金牌订单
	GoldMedalOrder bool `json:"gold_medal_order,omitempty" xml:"gold_medal_order,omitempty"`
	// 产品标识 student:学生票、nationality：国籍票、group：小团票、gold：金牌机票、card：宝贝机票、common：年龄票、speed：极速机票、age：限年龄票、delay：延时出票
	ProductTags []string `json:"product_tags,omitempty" xml:"product_tags>string,omitempty"`
	// 极速出票订单：true:是, false:否
	SpeedTicketOrder bool `json:"speed_ticket_order,omitempty" xml:"speed_ticket_order,omitempty"`
}
