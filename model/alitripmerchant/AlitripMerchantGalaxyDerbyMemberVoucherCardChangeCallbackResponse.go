package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 1
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackResponse.Put(v)
}
