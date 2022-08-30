package miniappopen

// TaobaoMiniappWidgetTemplateInstanceUpdateResult 结构体
type TaobaoMiniappWidgetTemplateInstanceUpdateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 返回结果
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 实体信息
	Model *MiniAppEntityTemplateDto `json:"model,omitempty" xml:"model,omitempty"`
}
