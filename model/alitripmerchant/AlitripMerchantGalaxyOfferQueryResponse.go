package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOfferQueryResponse 结构体
type AlitripMerchantGalaxyOfferQueryResponse struct {
	// offer列表
	Offers []OfferDetailsDto `json:"offers,omitempty" xml:"offers>offer_details_dto,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功还是失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyOfferQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOfferQueryResponse)
	},
}

// GetAlitripMerchantGalaxyOfferQueryResponse() 从对象池中获取AlitripMerchantGalaxyOfferQueryResponse
func GetAlitripMerchantGalaxyOfferQueryResponse() *AlitripMerchantGalaxyOfferQueryResponse {
	return poolAlitripMerchantGalaxyOfferQueryResponse.Get().(*AlitripMerchantGalaxyOfferQueryResponse)
}

// ReleaseAlitripMerchantGalaxyOfferQueryResponse 释放AlitripMerchantGalaxyOfferQueryResponse
func ReleaseAlitripMerchantGalaxyOfferQueryResponse(v *AlitripMerchantGalaxyOfferQueryResponse) {
	v.Offers = v.Offers[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyOfferQueryResponse.Put(v)
}
