package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMemberAddAgreementResponse 结构体
type AlitripMerchantGalaxyMemberAddAgreementResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyMemberAddAgreementResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMemberAddAgreementResponse)
	},
}

// GetAlitripMerchantGalaxyMemberAddAgreementResponse() 从对象池中获取AlitripMerchantGalaxyMemberAddAgreementResponse
func GetAlitripMerchantGalaxyMemberAddAgreementResponse() *AlitripMerchantGalaxyMemberAddAgreementResponse {
	return poolAlitripMerchantGalaxyMemberAddAgreementResponse.Get().(*AlitripMerchantGalaxyMemberAddAgreementResponse)
}

// ReleaseAlitripMerchantGalaxyMemberAddAgreementResponse 释放AlitripMerchantGalaxyMemberAddAgreementResponse
func ReleaseAlitripMerchantGalaxyMemberAddAgreementResponse(v *AlitripMerchantGalaxyMemberAddAgreementResponse) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyMemberAddAgreementResponse.Put(v)
}
