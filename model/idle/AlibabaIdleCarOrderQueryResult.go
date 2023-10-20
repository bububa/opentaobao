package idle

// AlibabaidlecarorderqueryResult 结构体
type AlibabaidlecarorderqueryResult struct {
	// 错误信息
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回结果
	Module *ConsignmentV2orderTo `json:"module,omitempty" xml:"module,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
