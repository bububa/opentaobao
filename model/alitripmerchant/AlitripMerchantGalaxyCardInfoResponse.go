package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyCardInfoResponse 结构体
type AlitripMerchantGalaxyCardInfoResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回内容
	Content *CardSystemVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyCardInfoResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCardInfoResponse)
	},
}

// GetAlitripMerchantGalaxyCardInfoResponse() 从对象池中获取AlitripMerchantGalaxyCardInfoResponse
func GetAlitripMerchantGalaxyCardInfoResponse() *AlitripMerchantGalaxyCardInfoResponse {
	return poolAlitripMerchantGalaxyCardInfoResponse.Get().(*AlitripMerchantGalaxyCardInfoResponse)
}

// ReleaseAlitripMerchantGalaxyCardInfoResponse 释放AlitripMerchantGalaxyCardInfoResponse
func ReleaseAlitripMerchantGalaxyCardInfoResponse(v *AlitripMerchantGalaxyCardInfoResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyCardInfoResponse.Put(v)
}
