package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 响应体
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundResponse.Put(v)
}
