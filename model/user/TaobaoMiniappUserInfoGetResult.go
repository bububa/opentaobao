package user

// TaobaoMiniappUserInfoGetResult 结构体
type TaobaoMiniappUserInfoGetResult struct {
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// model
	Model *OpenUserInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
