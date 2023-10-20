package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 内容
	Content *DerbyVoucherReceiptApplyVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsApplyResponse.Put(v)
}
