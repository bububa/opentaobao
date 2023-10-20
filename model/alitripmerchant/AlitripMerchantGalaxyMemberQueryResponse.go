package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMemberQueryResponse 结构体
type AlitripMerchantGalaxyMemberQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 用户信息
	Content *MemberDto `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyMemberQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberQueryResponse)
	},
}

// GetAlitripMerchantGalaxyMemberQueryResponse() 从对象池中获取AlitripMerchantGalaxyMemberQueryResponse
func GetAlitripMerchantGalaxyMemberQueryResponse() *AlitripMerchantGalaxyMemberQueryResponse {
	return poolAlitripMerchantGalaxyMemberQueryResponse.Get().(*AlitripMerchantGalaxyMemberQueryResponse)
}

// ReleaseAlitripMerchantGalaxyMemberQueryResponse 释放AlitripMerchantGalaxyMemberQueryResponse
func ReleaseAlitripMerchantGalaxyMemberQueryResponse(v *AlitripMerchantGalaxyMemberQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyMemberQueryResponse.Put(v)
}
