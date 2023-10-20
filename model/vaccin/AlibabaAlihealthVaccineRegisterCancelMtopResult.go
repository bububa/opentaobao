package vaccin

// AlibabaAlihealthVaccineRegisterCancelMtopResult 结构体
type AlibabaAlihealthVaccineRegisterCancelMtopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误提示
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 是否取消成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
