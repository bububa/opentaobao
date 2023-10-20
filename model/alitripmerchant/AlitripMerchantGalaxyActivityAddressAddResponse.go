package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyActivityAddressAddResponse 结构体
type AlitripMerchantGalaxyActivityAddressAddResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 邮寄信息是否填写成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyActivityAddressAddResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityAddressAddResponse)
	},
}

// GetAlitripMerchantGalaxyActivityAddressAddResponse() 从对象池中获取AlitripMerchantGalaxyActivityAddressAddResponse
func GetAlitripMerchantGalaxyActivityAddressAddResponse() *AlitripMerchantGalaxyActivityAddressAddResponse {
	return poolAlitripMerchantGalaxyActivityAddressAddResponse.Get().(*AlitripMerchantGalaxyActivityAddressAddResponse)
}

// ReleaseAlitripMerchantGalaxyActivityAddressAddResponse 释放AlitripMerchantGalaxyActivityAddressAddResponse
func ReleaseAlitripMerchantGalaxyActivityAddressAddResponse(v *AlitripMerchantGalaxyActivityAddressAddResponse) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Content = false
	v.Success = false
	poolAlitripMerchantGalaxyActivityAddressAddResponse.Put(v)
}
