package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMemberCompleteSwitchResponse 结构体
type AlitripMerchantGalaxyMemberCompleteSwitchResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否切换成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyMemberCompleteSwitchResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberCompleteSwitchResponse)
	},
}

// GetAlitripMerchantGalaxyMemberCompleteSwitchResponse() 从对象池中获取AlitripMerchantGalaxyMemberCompleteSwitchResponse
func GetAlitripMerchantGalaxyMemberCompleteSwitchResponse() *AlitripMerchantGalaxyMemberCompleteSwitchResponse {
	return poolAlitripMerchantGalaxyMemberCompleteSwitchResponse.Get().(*AlitripMerchantGalaxyMemberCompleteSwitchResponse)
}

// ReleaseAlitripMerchantGalaxyMemberCompleteSwitchResponse 释放AlitripMerchantGalaxyMemberCompleteSwitchResponse
func ReleaseAlitripMerchantGalaxyMemberCompleteSwitchResponse(v *AlitripMerchantGalaxyMemberCompleteSwitchResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyMemberCompleteSwitchResponse.Put(v)
}
