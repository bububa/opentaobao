package alitripmerchant

// AlitripMerchantGalaxyOrderCancelResponse 结构体
type AlitripMerchantGalaxyOrderCancelResponse struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果描述
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果成功判断
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
