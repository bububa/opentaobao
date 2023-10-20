package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse struct {
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 内容
	Content *DerbyVoucherReceiptQueryVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse() *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryResponse.Put(v)
}
