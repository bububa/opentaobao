package alitripmerchant

// AlitripMerchantGalaxyActivityDrawParticipateResponse 结构体
type AlitripMerchantGalaxyActivityDrawParticipateResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回类型
	Content *ActivityParticipateVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
