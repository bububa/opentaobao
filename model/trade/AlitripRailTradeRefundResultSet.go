package trade

// AlitriprailtraderefundResultSet 结构体
type AlitriprailtraderefundResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 接口成功/失败
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
