package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMemberRegisterDerbyResponse 结构体
type AlitripMerchantGalaxyMemberRegisterDerbyResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 试单结果
	Content *RegisterProgressVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyMemberRegisterDerbyResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberRegisterDerbyResponse)
	},
}

// GetAlitripMerchantGalaxyMemberRegisterDerbyResponse() 从对象池中获取AlitripMerchantGalaxyMemberRegisterDerbyResponse
func GetAlitripMerchantGalaxyMemberRegisterDerbyResponse() *AlitripMerchantGalaxyMemberRegisterDerbyResponse {
	return poolAlitripMerchantGalaxyMemberRegisterDerbyResponse.Get().(*AlitripMerchantGalaxyMemberRegisterDerbyResponse)
}

// ReleaseAlitripMerchantGalaxyMemberRegisterDerbyResponse 释放AlitripMerchantGalaxyMemberRegisterDerbyResponse
func ReleaseAlitripMerchantGalaxyMemberRegisterDerbyResponse(v *AlitripMerchantGalaxyMemberRegisterDerbyResponse) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyMemberRegisterDerbyResponse.Put(v)
}
