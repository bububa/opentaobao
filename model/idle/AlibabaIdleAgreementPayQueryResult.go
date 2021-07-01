package idle

// AlibabaIdleAgreementPayQueryResult 结构体
type AlibabaIdleAgreementPayQueryResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// err_code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 账单查询结果
	Module *Serializable `json:"module,omitempty" xml:"module,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
