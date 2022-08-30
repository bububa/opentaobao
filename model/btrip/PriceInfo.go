package btrip

// PriceInfo 结构体
type PriceInfo struct {
	// 乘机人，多个用‘,’分割
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 流水创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单交易流水号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 结算方式:1：个人现付，2:企业现付,4:企业月结，8、企业预存
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 价格
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
	// 交易类目编码
	CategoryCode int64 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// 资金流向:1:支出，2:收入
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
