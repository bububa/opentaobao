package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMemberTokenResponse 结构体
type AlitripMerchantGalaxyMemberTokenResponse struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功还是失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyMemberTokenResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberTokenResponse)
	},
}

// GetAlitripMerchantGalaxyMemberTokenResponse() 从对象池中获取AlitripMerchantGalaxyMemberTokenResponse
func GetAlitripMerchantGalaxyMemberTokenResponse() *AlitripMerchantGalaxyMemberTokenResponse {
	return poolAlitripMerchantGalaxyMemberTokenResponse.Get().(*AlitripMerchantGalaxyMemberTokenResponse)
}

// ReleaseAlitripMerchantGalaxyMemberTokenResponse 释放AlitripMerchantGalaxyMemberTokenResponse
func ReleaseAlitripMerchantGalaxyMemberTokenResponse(v *AlitripMerchantGalaxyMemberTokenResponse) {
	v.ErrorCode = ""
	v.Token = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyMemberTokenResponse.Put(v)
}
