package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatLoginResponse 结构体
type AlitripMerchantGalaxyWechatLoginResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyWechatLoginResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatLoginResponse)
	},
}

// GetAlitripMerchantGalaxyWechatLoginResponse() 从对象池中获取AlitripMerchantGalaxyWechatLoginResponse
func GetAlitripMerchantGalaxyWechatLoginResponse() *AlitripMerchantGalaxyWechatLoginResponse {
	return poolAlitripMerchantGalaxyWechatLoginResponse.Get().(*AlitripMerchantGalaxyWechatLoginResponse)
}

// ReleaseAlitripMerchantGalaxyWechatLoginResponse 释放AlitripMerchantGalaxyWechatLoginResponse
func ReleaseAlitripMerchantGalaxyWechatLoginResponse(v *AlitripMerchantGalaxyWechatLoginResponse) {
	v.ErrorCode = ""
	v.Token = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyWechatLoginResponse.Put(v)
}
