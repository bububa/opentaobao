package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单取消成功
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = ""
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelResponse.Put(v)
}
