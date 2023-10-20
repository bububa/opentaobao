package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse struct {
	// 代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功/失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse() *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusResponse.Put(v)
}
