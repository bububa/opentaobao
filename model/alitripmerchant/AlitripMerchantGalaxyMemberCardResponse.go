package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMemberCardResponse 结构体
type AlitripMerchantGalaxyMemberCardResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回参数
	Content *MemberCardDto `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyMemberCardResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberCardResponse)
	},
}

// GetAlitripMerchantGalaxyMemberCardResponse() 从对象池中获取AlitripMerchantGalaxyMemberCardResponse
func GetAlitripMerchantGalaxyMemberCardResponse() *AlitripMerchantGalaxyMemberCardResponse {
	return poolAlitripMerchantGalaxyMemberCardResponse.Get().(*AlitripMerchantGalaxyMemberCardResponse)
}

// ReleaseAlitripMerchantGalaxyMemberCardResponse 释放AlitripMerchantGalaxyMemberCardResponse
func ReleaseAlitripMerchantGalaxyMemberCardResponse(v *AlitripMerchantGalaxyMemberCardResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyMemberCardResponse.Put(v)
}
