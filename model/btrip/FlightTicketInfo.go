package btrip

// FlightTicketInfo 结构体
type FlightTicketInfo struct {
	// 更新时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 票状态
	TicketStatus string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	// 票价
	TicketPrice float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 折扣
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 燃油费用
	OilPrice float64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// 支付方式 1个人支付/2企业支付/3混付
	PayType float64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 机建费用
	BuildPrice float64 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// 票状态编码
	TicketStatusCode int64 `json:"ticket_status_code,omitempty" xml:"ticket_status_code,omitempty"`
	// 是否改签
	Changed bool `json:"changed,omitempty" xml:"changed,omitempty"`
}
