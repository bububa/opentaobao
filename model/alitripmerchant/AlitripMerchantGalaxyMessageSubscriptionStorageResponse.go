package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyMessageSubscriptionStorageResponse 结构体
type AlitripMerchantGalaxyMessageSubscriptionStorageResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 信息
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyMessageSubscriptionStorageResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyMessageSubscriptionStorageResponse)
	},
}

// GetAlitripMerchantGalaxyMessageSubscriptionStorageResponse() 从对象池中获取AlitripMerchantGalaxyMessageSubscriptionStorageResponse
func GetAlitripMerchantGalaxyMessageSubscriptionStorageResponse() *AlitripMerchantGalaxyMessageSubscriptionStorageResponse {
	return poolAlitripMerchantGalaxyMessageSubscriptionStorageResponse.Get().(*AlitripMerchantGalaxyMessageSubscriptionStorageResponse)
}

// ReleaseAlitripMerchantGalaxyMessageSubscriptionStorageResponse 释放AlitripMerchantGalaxyMessageSubscriptionStorageResponse
func ReleaseAlitripMerchantGalaxyMessageSubscriptionStorageResponse(v *AlitripMerchantGalaxyMessageSubscriptionStorageResponse) {
	v.ErrorCode = ""
	v.Content = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyMessageSubscriptionStorageResponse.Put(v)
}
