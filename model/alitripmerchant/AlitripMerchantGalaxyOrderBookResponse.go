package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOrderBookResponse 结构体
type AlitripMerchantGalaxyOrderBookResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 预订返回
	Content *PayParamResult `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyOrderBookResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderBookResponse)
	},
}

// GetAlitripMerchantGalaxyOrderBookResponse() 从对象池中获取AlitripMerchantGalaxyOrderBookResponse
func GetAlitripMerchantGalaxyOrderBookResponse() *AlitripMerchantGalaxyOrderBookResponse {
	return poolAlitripMerchantGalaxyOrderBookResponse.Get().(*AlitripMerchantGalaxyOrderBookResponse)
}

// ReleaseAlitripMerchantGalaxyOrderBookResponse 释放AlitripMerchantGalaxyOrderBookResponse
func ReleaseAlitripMerchantGalaxyOrderBookResponse(v *AlitripMerchantGalaxyOrderBookResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyOrderBookResponse.Put(v)
}
