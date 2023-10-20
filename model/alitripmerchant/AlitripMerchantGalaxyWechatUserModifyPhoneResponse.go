package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatUserModifyPhoneResponse 结构体
type AlitripMerchantGalaxyWechatUserModifyPhoneResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 用户token及用户当前状态
	Content *UserCurrentStatus `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyWechatUserModifyPhoneResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatUserModifyPhoneResponse)
	},
}

// GetAlitripMerchantGalaxyWechatUserModifyPhoneResponse() 从对象池中获取AlitripMerchantGalaxyWechatUserModifyPhoneResponse
func GetAlitripMerchantGalaxyWechatUserModifyPhoneResponse() *AlitripMerchantGalaxyWechatUserModifyPhoneResponse {
	return poolAlitripMerchantGalaxyWechatUserModifyPhoneResponse.Get().(*AlitripMerchantGalaxyWechatUserModifyPhoneResponse)
}

// ReleaseAlitripMerchantGalaxyWechatUserModifyPhoneResponse 释放AlitripMerchantGalaxyWechatUserModifyPhoneResponse
func ReleaseAlitripMerchantGalaxyWechatUserModifyPhoneResponse(v *AlitripMerchantGalaxyWechatUserModifyPhoneResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyWechatUserModifyPhoneResponse.Put(v)
}
