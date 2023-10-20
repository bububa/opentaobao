package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatAddOperationRecordResponse 结构体
type AlitripMerchantGalaxyWechatAddOperationRecordResponse struct {
	// 状态码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 消息体
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlitripMerchantGalaxyWechatAddOperationRecordResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatAddOperationRecordResponse)
	},
}

// GetAlitripMerchantGalaxyWechatAddOperationRecordResponse() 从对象池中获取AlitripMerchantGalaxyWechatAddOperationRecordResponse
func GetAlitripMerchantGalaxyWechatAddOperationRecordResponse() *AlitripMerchantGalaxyWechatAddOperationRecordResponse {
	return poolAlitripMerchantGalaxyWechatAddOperationRecordResponse.Get().(*AlitripMerchantGalaxyWechatAddOperationRecordResponse)
}

// ReleaseAlitripMerchantGalaxyWechatAddOperationRecordResponse 释放AlitripMerchantGalaxyWechatAddOperationRecordResponse
func ReleaseAlitripMerchantGalaxyWechatAddOperationRecordResponse(v *AlitripMerchantGalaxyWechatAddOperationRecordResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Content = false
	poolAlitripMerchantGalaxyWechatAddOperationRecordResponse.Put(v)
}
