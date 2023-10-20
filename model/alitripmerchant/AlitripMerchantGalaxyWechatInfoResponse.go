package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatInfoResponse 结构体
type AlitripMerchantGalaxyWechatInfoResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回的微信用户信息
	Content *WechatCodeResponse `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyWechatInfoResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatInfoResponse)
	},
}

// GetAlitripMerchantGalaxyWechatInfoResponse() 从对象池中获取AlitripMerchantGalaxyWechatInfoResponse
func GetAlitripMerchantGalaxyWechatInfoResponse() *AlitripMerchantGalaxyWechatInfoResponse {
	return poolAlitripMerchantGalaxyWechatInfoResponse.Get().(*AlitripMerchantGalaxyWechatInfoResponse)
}

// ReleaseAlitripMerchantGalaxyWechatInfoResponse 释放AlitripMerchantGalaxyWechatInfoResponse
func ReleaseAlitripMerchantGalaxyWechatInfoResponse(v *AlitripMerchantGalaxyWechatInfoResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyWechatInfoResponse.Put(v)
}
