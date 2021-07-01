package alitripmerchant

// AlitripMerchantGalaxyOfferQueryResponse 结构体
type AlitripMerchantGalaxyOfferQueryResponse struct {
	// 成功还是失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// offer列表
	Offers []OfferDetailsDto `json:"offers,omitempty" xml:"offers>offer_details_dto,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
