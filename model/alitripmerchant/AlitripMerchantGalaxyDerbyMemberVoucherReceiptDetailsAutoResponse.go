package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse struct {
	// 返回类型
	Content []DerbyVoucherReceiptAutoVo `json:"content,omitempty" xml:"content>derby_voucher_receipt_auto_vo,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse) {
	v.Content = v.Content[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoResponse.Put(v)
}
