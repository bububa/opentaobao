package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOrderQueryResponse 结构体
type AlitripMerchantGalaxyOrderQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单详情
	Content *OrderDetailDto `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyOrderQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderQueryResponse)
	},
}

// GetAlitripMerchantGalaxyOrderQueryResponse() 从对象池中获取AlitripMerchantGalaxyOrderQueryResponse
func GetAlitripMerchantGalaxyOrderQueryResponse() *AlitripMerchantGalaxyOrderQueryResponse {
	return poolAlitripMerchantGalaxyOrderQueryResponse.Get().(*AlitripMerchantGalaxyOrderQueryResponse)
}

// ReleaseAlitripMerchantGalaxyOrderQueryResponse 释放AlitripMerchantGalaxyOrderQueryResponse
func ReleaseAlitripMerchantGalaxyOrderQueryResponse(v *AlitripMerchantGalaxyOrderQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyOrderQueryResponse.Put(v)
}
