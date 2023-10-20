package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOrderFillResponse 结构体
type AlitripMerchantGalaxyOrderFillResponse struct {
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 填单页展示对象
	Content *FillPageVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyOrderFillResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderFillResponse)
	},
}

// GetAlitripMerchantGalaxyOrderFillResponse() 从对象池中获取AlitripMerchantGalaxyOrderFillResponse
func GetAlitripMerchantGalaxyOrderFillResponse() *AlitripMerchantGalaxyOrderFillResponse {
	return poolAlitripMerchantGalaxyOrderFillResponse.Get().(*AlitripMerchantGalaxyOrderFillResponse)
}

// ReleaseAlitripMerchantGalaxyOrderFillResponse 释放AlitripMerchantGalaxyOrderFillResponse
func ReleaseAlitripMerchantGalaxyOrderFillResponse(v *AlitripMerchantGalaxyOrderFillResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyOrderFillResponse.Put(v)
}
