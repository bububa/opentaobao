package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOrderValidateResponse 结构体
type AlitripMerchantGalaxyOrderValidateResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 试单结果
	Content *ValidateOrderVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyOrderValidateResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderValidateResponse)
	},
}

// GetAlitripMerchantGalaxyOrderValidateResponse() 从对象池中获取AlitripMerchantGalaxyOrderValidateResponse
func GetAlitripMerchantGalaxyOrderValidateResponse() *AlitripMerchantGalaxyOrderValidateResponse {
	return poolAlitripMerchantGalaxyOrderValidateResponse.Get().(*AlitripMerchantGalaxyOrderValidateResponse)
}

// ReleaseAlitripMerchantGalaxyOrderValidateResponse 释放AlitripMerchantGalaxyOrderValidateResponse
func ReleaseAlitripMerchantGalaxyOrderValidateResponse(v *AlitripMerchantGalaxyOrderValidateResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyOrderValidateResponse.Put(v)
}
