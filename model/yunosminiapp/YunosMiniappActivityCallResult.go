package yunosminiapp

// YunosMiniappActivityCallResult 结构体
type YunosMiniappActivityCallResult struct {
	// 返回信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回内容
	Result *TopActivityResult `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
