package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse 结构体
type AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回实体内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse
func GetAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse() *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse {
	return poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse.Get().(*AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse 释放AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse(v *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = ""
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse.Put(v)
}
