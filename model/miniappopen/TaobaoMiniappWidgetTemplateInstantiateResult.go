package miniappopen

// TaobaominiappwidgettemplateinstantiateResult 结构体
type TaobaominiappwidgettemplateinstantiateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 返回实体信息
	Model *MiniAppEntityTemplateDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
