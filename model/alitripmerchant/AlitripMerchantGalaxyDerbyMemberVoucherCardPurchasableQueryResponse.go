package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 可购权益卡实体类
	Content *DerbyVoucherCardPurchasableVo `json:"content,omitempty" xml:"content,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse.Put(v)
}
