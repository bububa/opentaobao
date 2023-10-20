package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatPayCallbackResponse 结构体
type AlitripMerchantGalaxyWechatPayCallbackResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 微信要求的返回类
	Content *WechatCallbackResponse `json:"content,omitempty" xml:"content,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyWechatPayCallbackResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatPayCallbackResponse)
	},
}

// GetAlitripMerchantGalaxyWechatPayCallbackResponse() 从对象池中获取AlitripMerchantGalaxyWechatPayCallbackResponse
func GetAlitripMerchantGalaxyWechatPayCallbackResponse() *AlitripMerchantGalaxyWechatPayCallbackResponse {
	return poolAlitripMerchantGalaxyWechatPayCallbackResponse.Get().(*AlitripMerchantGalaxyWechatPayCallbackResponse)
}

// ReleaseAlitripMerchantGalaxyWechatPayCallbackResponse 释放AlitripMerchantGalaxyWechatPayCallbackResponse
func ReleaseAlitripMerchantGalaxyWechatPayCallbackResponse(v *AlitripMerchantGalaxyWechatPayCallbackResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyWechatPayCallbackResponse.Put(v)
}
