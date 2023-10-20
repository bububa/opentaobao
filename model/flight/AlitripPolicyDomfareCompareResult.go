package flight

// AlitrippolicydomfarecompareResult 结构体
type AlitrippolicydomfarecompareResult struct {
	// 调用错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码详情
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 比价结果列表
	Data *CompareDomFareReponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
