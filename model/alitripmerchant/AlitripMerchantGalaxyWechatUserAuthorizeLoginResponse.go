package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse 结构体
type AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 具体错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 用户token相关信息
	Content *UserCurrentStatus `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyWechatUserAuthorizeLoginResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse)
	},
}

// GetAlitripMerchantGalaxyWechatUserAuthorizeLoginResponse() 从对象池中获取AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse
func GetAlitripMerchantGalaxyWechatUserAuthorizeLoginResponse() *AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse {
	return poolAlitripMerchantGalaxyWechatUserAuthorizeLoginResponse.Get().(*AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse)
}

// ReleaseAlitripMerchantGalaxyWechatUserAuthorizeLoginResponse 释放AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse
func ReleaseAlitripMerchantGalaxyWechatUserAuthorizeLoginResponse(v *AlitripMerchantGalaxyWechatUserAuthorizeLoginResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyWechatUserAuthorizeLoginResponse.Put(v)
}
