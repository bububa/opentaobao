package yunosminiapp

// YunosMiniappDatatunnelCallMapResult 结构体
type YunosMiniappDatatunnelCallMapResult struct {
	// 随机字符串
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果核心内容
	Result *YunosMiniappDatatunnelCallResult `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
