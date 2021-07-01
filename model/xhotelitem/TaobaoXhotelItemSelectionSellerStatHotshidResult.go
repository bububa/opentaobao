package xhotelitem

// TaobaoXhotelItemSelectionSellerStatHotshidResult 结构体
type TaobaoXhotelItemSelectionSellerStatHotshidResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回结构
	Module *TaobaoXhotelItemSelectionSellerStatHotshidModule `json:"module,omitempty" xml:"module,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
