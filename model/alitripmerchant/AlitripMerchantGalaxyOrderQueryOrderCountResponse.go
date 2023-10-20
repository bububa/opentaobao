package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOrderQueryOrderCountResponse 结构体
type AlitripMerchantGalaxyOrderQueryOrderCountResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回实体
	Content *OrderCountVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyOrderQueryOrderCountResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderQueryOrderCountResponse)
	},
}

// GetAlitripMerchantGalaxyOrderQueryOrderCountResponse() 从对象池中获取AlitripMerchantGalaxyOrderQueryOrderCountResponse
func GetAlitripMerchantGalaxyOrderQueryOrderCountResponse() *AlitripMerchantGalaxyOrderQueryOrderCountResponse {
	return poolAlitripMerchantGalaxyOrderQueryOrderCountResponse.Get().(*AlitripMerchantGalaxyOrderQueryOrderCountResponse)
}

// ReleaseAlitripMerchantGalaxyOrderQueryOrderCountResponse 释放AlitripMerchantGalaxyOrderQueryOrderCountResponse
func ReleaseAlitripMerchantGalaxyOrderQueryOrderCountResponse(v *AlitripMerchantGalaxyOrderQueryOrderCountResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyOrderQueryOrderCountResponse.Put(v)
}
