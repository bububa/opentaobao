package c2m

// TaobaoTxpItemItemlistgetResultDto 结构体
type TaobaoTxpItemItemlistgetResultDto struct {
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回的结果信息
	Model *ItemTopResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 执行的错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
