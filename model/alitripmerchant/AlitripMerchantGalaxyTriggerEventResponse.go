package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyTriggerEventResponse 结构体
type AlitripMerchantGalaxyTriggerEventResponse struct {
	// 状态码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 状态信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 消息体
	Content *EventCoupon `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyTriggerEventResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyTriggerEventResponse)
	},
}

// GetAlitripMerchantGalaxyTriggerEventResponse() 从对象池中获取AlitripMerchantGalaxyTriggerEventResponse
func GetAlitripMerchantGalaxyTriggerEventResponse() *AlitripMerchantGalaxyTriggerEventResponse {
	return poolAlitripMerchantGalaxyTriggerEventResponse.Get().(*AlitripMerchantGalaxyTriggerEventResponse)
}

// ReleaseAlitripMerchantGalaxyTriggerEventResponse 释放AlitripMerchantGalaxyTriggerEventResponse
func ReleaseAlitripMerchantGalaxyTriggerEventResponse(v *AlitripMerchantGalaxyTriggerEventResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyTriggerEventResponse.Put(v)
}
