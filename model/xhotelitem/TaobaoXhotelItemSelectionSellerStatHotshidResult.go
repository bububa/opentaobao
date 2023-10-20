package xhotelitem

// TaobaoxhotelitemselectionsellerstathotshidResult 结构体
type TaobaoxhotelitemselectionsellerstathotshidResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结构
	Module *TaobaoxhotelitemselectionsellerstathotshidModule `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
