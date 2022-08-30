package btrip

// OpenPriceInfo 结构体
type OpenPriceInfo struct {
	// 交易类目
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 流水创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 支付流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 乘机人，多个用‘,’分割
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 订单交易流水号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 改签前的票号
	OriginalTicketNo string `json:"original_ticket_no,omitempty" xml:"original_ticket_no,omitempty"`
	// 改签票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 改签航班号
	ChangeFlightNo string `json:"change_flight_no,omitempty" xml:"change_flight_no,omitempty"`
	// 改签折扣
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 改签票起飞时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 改签票到达时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 改签前原车次号
	OriginalTrainNo string `json:"original_train_no,omitempty" xml:"original_train_no,omitempty"`
	// 改签车次号
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 改签座席
	SeatType string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// 改签出发城市
	StartCity string `json:"start_city,omitempty" xml:"start_city,omitempty"`
	// 改签到达城市
	EndCity string `json:"end_city,omitempty" xml:"end_city,omitempty"`
	// 个人支付金额
	PersonPrice string `json:"person_price,omitempty" xml:"person_price,omitempty"`
	// 结算方式:1：个人现付，2:企业现付,4:企业月结，8、企业预存
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 资金流向:1:支出，2:收入
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 交易类目类型1-飞机，2-酒店，3-火车，4-用车
	CategoryType int64 `json:"category_type,omitempty" xml:"category_type,omitempty"`
	// 交易类目编码
	CategoryCode int64 `json:"category_code,omitempty" xml:"category_code,omitempty"`
}
