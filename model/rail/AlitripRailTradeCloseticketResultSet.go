package rail

// AlitripRailTradeCloseticketResultSet 结构体
type AlitripRailTradeCloseticketResultSet struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功失败
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
