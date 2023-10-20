package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 可购权益卡实体类
	Content *DerbyVoucherCardPurchasableVo `json:"content,omitempty" xml:"content,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderPlaceResponse.Put(v)
}
