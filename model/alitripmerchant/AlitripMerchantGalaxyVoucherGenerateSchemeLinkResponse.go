package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse 结构体
type AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse)
	},
}

// GetAlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse() 从对象池中获取AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse
func GetAlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse() *AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse {
	return poolAlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse.Get().(*AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse)
}

// ReleaseAlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse 释放AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse
func ReleaseAlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse(v *AlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse) {
	v.ErrorCode = ""
	v.Content = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyVoucherGenerateSchemeLinkResponse.Put(v)
}
