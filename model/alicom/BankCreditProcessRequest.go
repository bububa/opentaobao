package alicom

// BankCreditProcessRequest 结构体
type BankCreditProcessRequest struct {
	// 银行编码
	BankCode string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
	// 事件类型
	EventType string `json:"event_type,omitempty" xml:"event_type,omitempty"`
	// 用户的openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 事件时间时间戳
	EventTime int64 `json:"event_time,omitempty" xml:"event_time,omitempty"`
	// 是否该自然人首刷
	NewPay bool `json:"new_pay,omitempty" xml:"new_pay,omitempty"`
	// 是否该自然人首卡
	NewCustomer bool `json:"new_customer,omitempty" xml:"new_customer,omitempty"`
}
