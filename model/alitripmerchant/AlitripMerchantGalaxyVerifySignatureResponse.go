package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyVerifySignatureResponse 结构体
type AlitripMerchantGalaxyVerifySignatureResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回内容
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyVerifySignatureResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyVerifySignatureResponse)
	},
}

// GetAlitripMerchantGalaxyVerifySignatureResponse() 从对象池中获取AlitripMerchantGalaxyVerifySignatureResponse
func GetAlitripMerchantGalaxyVerifySignatureResponse() *AlitripMerchantGalaxyVerifySignatureResponse {
	return poolAlitripMerchantGalaxyVerifySignatureResponse.Get().(*AlitripMerchantGalaxyVerifySignatureResponse)
}

// ReleaseAlitripMerchantGalaxyVerifySignatureResponse 释放AlitripMerchantGalaxyVerifySignatureResponse
func ReleaseAlitripMerchantGalaxyVerifySignatureResponse(v *AlitripMerchantGalaxyVerifySignatureResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyVerifySignatureResponse.Put(v)
}
