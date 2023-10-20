package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMemberProviderRegisterResponse 结构体
type AlitripMerchantGalaxyMemberProviderRegisterResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否注册成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyMemberProviderRegisterResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberProviderRegisterResponse)
	},
}

// GetAlitripMerchantGalaxyMemberProviderRegisterResponse() 从对象池中获取AlitripMerchantGalaxyMemberProviderRegisterResponse
func GetAlitripMerchantGalaxyMemberProviderRegisterResponse() *AlitripMerchantGalaxyMemberProviderRegisterResponse {
	return poolAlitripMerchantGalaxyMemberProviderRegisterResponse.Get().(*AlitripMerchantGalaxyMemberProviderRegisterResponse)
}

// ReleaseAlitripMerchantGalaxyMemberProviderRegisterResponse 释放AlitripMerchantGalaxyMemberProviderRegisterResponse
func ReleaseAlitripMerchantGalaxyMemberProviderRegisterResponse(v *AlitripMerchantGalaxyMemberProviderRegisterResponse) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Content = false
	v.Success = false
	poolAlitripMerchantGalaxyMemberProviderRegisterResponse.Put(v)
}
