package alitripmerchant

// AlitripMerchantGalaxyQueryDrawSummaryResponse 结构体
type AlitripMerchantGalaxyQueryDrawSummaryResponse struct {
	// 返回类型
	Contents []DrawOfferSummaryVo `json:"contents,omitempty" xml:"contents>draw_offer_summary_vo,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
