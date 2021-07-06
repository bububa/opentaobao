package ascm

// SettlementGatewayResult 结构体
type SettlementGatewayResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
