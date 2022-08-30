package alitripmerchant

// AlitripMerchantGalaxyActivityFatigueResponse 结构体
type AlitripMerchantGalaxyActivityFatigueResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回数据
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
