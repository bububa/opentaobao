package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 500
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1com.fliggy.merchant.unicorn.api.derbyVoucherCard.vo.DerbyVoucherIdentityQrcodeVO
	Content *DerbyVoucherIdentityQrcodeVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse.Put(v)
}
