package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse struct {
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse) {
	v.Success = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = ""
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptShowResponse.Put(v)
}
