package alitripmerchant

// AlitripMerchantGalaxyCityLikeResponse 结构体
type AlitripMerchantGalaxyCityLikeResponse struct {
	// 返回结果
	Contents []AddressSearchDto `json:"contents,omitempty" xml:"contents>address_search_dto,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
