package exchange

// RefundBaseResponse 结构体
type RefundBaseResponse struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// exchange
	Exchange *Exchange `json:"exchange,omitempty" xml:"exchange,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
