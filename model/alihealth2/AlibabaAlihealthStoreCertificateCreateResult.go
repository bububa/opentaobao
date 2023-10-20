package alihealth2

// AlibabaalihealthstorecertificatecreateResult 结构体
type AlibabaalihealthstorecertificatecreateResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
	// 创建审批流是否成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
