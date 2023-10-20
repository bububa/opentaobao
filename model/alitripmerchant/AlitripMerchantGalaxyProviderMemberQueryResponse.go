package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyProviderMemberQueryResponse 结构体
type AlitripMerchantGalaxyProviderMemberQueryResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果
	ProviderMemberVo *ProviderMemberVo `json:"provider_member_vo,omitempty" xml:"provider_member_vo,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyProviderMemberQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyProviderMemberQueryResponse)
	},
}

// GetAlitripMerchantGalaxyProviderMemberQueryResponse() 从对象池中获取AlitripMerchantGalaxyProviderMemberQueryResponse
func GetAlitripMerchantGalaxyProviderMemberQueryResponse() *AlitripMerchantGalaxyProviderMemberQueryResponse {
	return poolAlitripMerchantGalaxyProviderMemberQueryResponse.Get().(*AlitripMerchantGalaxyProviderMemberQueryResponse)
}

// ReleaseAlitripMerchantGalaxyProviderMemberQueryResponse 释放AlitripMerchantGalaxyProviderMemberQueryResponse
func ReleaseAlitripMerchantGalaxyProviderMemberQueryResponse(v *AlitripMerchantGalaxyProviderMemberQueryResponse) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.ProviderMemberVo = nil
	v.Success = false
	poolAlitripMerchantGalaxyProviderMemberQueryResponse.Put(v)
}
