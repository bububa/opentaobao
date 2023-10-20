package idle

// AlibabaidlerecycleordershowResult 结构体
type AlibabaidlerecycleordershowResult struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 订单信息
	Module *Serializable `json:"module,omitempty" xml:"module,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
