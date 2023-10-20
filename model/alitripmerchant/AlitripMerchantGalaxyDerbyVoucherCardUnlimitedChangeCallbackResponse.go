package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse 结构体
type AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 回调处理结果
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse() 从对象池中获取AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse
func GetAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse() *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse {
	return poolAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse.Get().(*AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse 释放AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse
func ReleaseAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse(v *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse.Put(v)
}
