package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 内容
	Content *DerbyVoucherReceiptAutoVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoResponse.Put(v)
}
