package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatCardQueryRecordResponse 结构体
type AlitripMerchantGalaxyWechatCardQueryRecordResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyWechatCardQueryRecordResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatCardQueryRecordResponse)
	},
}

// GetAlitripMerchantGalaxyWechatCardQueryRecordResponse() 从对象池中获取AlitripMerchantGalaxyWechatCardQueryRecordResponse
func GetAlitripMerchantGalaxyWechatCardQueryRecordResponse() *AlitripMerchantGalaxyWechatCardQueryRecordResponse {
	return poolAlitripMerchantGalaxyWechatCardQueryRecordResponse.Get().(*AlitripMerchantGalaxyWechatCardQueryRecordResponse)
}

// ReleaseAlitripMerchantGalaxyWechatCardQueryRecordResponse 释放AlitripMerchantGalaxyWechatCardQueryRecordResponse
func ReleaseAlitripMerchantGalaxyWechatCardQueryRecordResponse(v *AlitripMerchantGalaxyWechatCardQueryRecordResponse) {
	v.ErrorCode = ""
	v.Content = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyWechatCardQueryRecordResponse.Put(v)
}
