package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMemberLogoutResponse 结构体
type AlitripMerchantGalaxyMemberLogoutResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 登出是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlitripMerchantGalaxyMemberLogoutResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberLogoutResponse)
	},
}

// GetAlitripMerchantGalaxyMemberLogoutResponse() 从对象池中获取AlitripMerchantGalaxyMemberLogoutResponse
func GetAlitripMerchantGalaxyMemberLogoutResponse() *AlitripMerchantGalaxyMemberLogoutResponse {
	return poolAlitripMerchantGalaxyMemberLogoutResponse.Get().(*AlitripMerchantGalaxyMemberLogoutResponse)
}

// ReleaseAlitripMerchantGalaxyMemberLogoutResponse 释放AlitripMerchantGalaxyMemberLogoutResponse
func ReleaseAlitripMerchantGalaxyMemberLogoutResponse(v *AlitripMerchantGalaxyMemberLogoutResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.IsSuccess = false
	poolAlitripMerchantGalaxyMemberLogoutResponse.Put(v)
}
