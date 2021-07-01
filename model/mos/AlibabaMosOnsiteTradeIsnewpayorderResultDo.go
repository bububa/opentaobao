package mos

// AlibabaMosOnsiteTradeIsnewpayorderResultDo 结构体
type AlibabaMosOnsiteTradeIsnewpayorderResultDo struct {
	// 是否为新支付订单。true：是，false:不是
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// errCode
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
