package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatUserLoginResponse 结构体
type AlitripMerchantGalaxyWechatUserLoginResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 用户token及用户当前状态
	Content *UserCurrentStatus `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyWechatUserLoginResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatUserLoginResponse)
	},
}

// GetAlitripMerchantGalaxyWechatUserLoginResponse() 从对象池中获取AlitripMerchantGalaxyWechatUserLoginResponse
func GetAlitripMerchantGalaxyWechatUserLoginResponse() *AlitripMerchantGalaxyWechatUserLoginResponse {
	return poolAlitripMerchantGalaxyWechatUserLoginResponse.Get().(*AlitripMerchantGalaxyWechatUserLoginResponse)
}

// ReleaseAlitripMerchantGalaxyWechatUserLoginResponse 释放AlitripMerchantGalaxyWechatUserLoginResponse
func ReleaseAlitripMerchantGalaxyWechatUserLoginResponse(v *AlitripMerchantGalaxyWechatUserLoginResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyWechatUserLoginResponse.Put(v)
}
