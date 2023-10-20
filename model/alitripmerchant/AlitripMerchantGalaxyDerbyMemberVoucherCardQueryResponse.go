package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse struct {
	// 1
	Content []DerbyVoucherCardVo `json:"content,omitempty" xml:"content>derby_voucher_card_vo,omitempty"`
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse) {
	v.Content = v.Content[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse.Put(v)
}
