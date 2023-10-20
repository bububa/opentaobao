package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyActivityPopupControlResponse 结构体
type AlitripMerchantGalaxyActivityPopupControlResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否执行成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyActivityPopupControlResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityPopupControlResponse)
	},
}

// GetAlitripMerchantGalaxyActivityPopupControlResponse() 从对象池中获取AlitripMerchantGalaxyActivityPopupControlResponse
func GetAlitripMerchantGalaxyActivityPopupControlResponse() *AlitripMerchantGalaxyActivityPopupControlResponse {
	return poolAlitripMerchantGalaxyActivityPopupControlResponse.Get().(*AlitripMerchantGalaxyActivityPopupControlResponse)
}

// ReleaseAlitripMerchantGalaxyActivityPopupControlResponse 释放AlitripMerchantGalaxyActivityPopupControlResponse
func ReleaseAlitripMerchantGalaxyActivityPopupControlResponse(v *AlitripMerchantGalaxyActivityPopupControlResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyActivityPopupControlResponse.Put(v)
}
