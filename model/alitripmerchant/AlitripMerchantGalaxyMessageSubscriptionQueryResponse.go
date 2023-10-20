package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMessageSubscriptionQueryResponse 结构体
type AlitripMerchantGalaxyMessageSubscriptionQueryResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 权限是否存在
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyMessageSubscriptionQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMessageSubscriptionQueryResponse)
	},
}

// GetAlitripMerchantGalaxyMessageSubscriptionQueryResponse() 从对象池中获取AlitripMerchantGalaxyMessageSubscriptionQueryResponse
func GetAlitripMerchantGalaxyMessageSubscriptionQueryResponse() *AlitripMerchantGalaxyMessageSubscriptionQueryResponse {
	return poolAlitripMerchantGalaxyMessageSubscriptionQueryResponse.Get().(*AlitripMerchantGalaxyMessageSubscriptionQueryResponse)
}

// ReleaseAlitripMerchantGalaxyMessageSubscriptionQueryResponse 释放AlitripMerchantGalaxyMessageSubscriptionQueryResponse
func ReleaseAlitripMerchantGalaxyMessageSubscriptionQueryResponse(v *AlitripMerchantGalaxyMessageSubscriptionQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyMessageSubscriptionQueryResponse.Put(v)
}
