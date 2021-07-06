package mos

// AlibabaMosOnsiteTradeQueryrefundResultDo 结构体
type AlibabaMosOnsiteTradeQueryrefundResultDo struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *OnsiteRefundResponse `json:"data,omitempty" xml:"data,omitempty"`
	// errCode
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
