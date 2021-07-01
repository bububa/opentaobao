package alicom

// AlibabaAliqinTaSmsNumSendBizResult 结构体
type AlibabaAliqinTaSmsNumSendBizResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}
