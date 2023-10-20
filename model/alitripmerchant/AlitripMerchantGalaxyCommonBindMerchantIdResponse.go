package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyCommonBindMerchantIdResponse 结构体
type AlitripMerchantGalaxyCommonBindMerchantIdResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyCommonBindMerchantIdResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCommonBindMerchantIdResponse)
	},
}

// GetAlitripMerchantGalaxyCommonBindMerchantIdResponse() 从对象池中获取AlitripMerchantGalaxyCommonBindMerchantIdResponse
func GetAlitripMerchantGalaxyCommonBindMerchantIdResponse() *AlitripMerchantGalaxyCommonBindMerchantIdResponse {
	return poolAlitripMerchantGalaxyCommonBindMerchantIdResponse.Get().(*AlitripMerchantGalaxyCommonBindMerchantIdResponse)
}

// ReleaseAlitripMerchantGalaxyCommonBindMerchantIdResponse 释放AlitripMerchantGalaxyCommonBindMerchantIdResponse
func ReleaseAlitripMerchantGalaxyCommonBindMerchantIdResponse(v *AlitripMerchantGalaxyCommonBindMerchantIdResponse) {
	v.ErrorCode = ""
	v.Content = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyCommonBindMerchantIdResponse.Put(v)
}
