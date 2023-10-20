package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyQueryParticipateNumberResponse 结构体
type AlitripMerchantGalaxyQueryParticipateNumberResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回类型
	Content *ActivityParticipateNumberVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyQueryParticipateNumberResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyQueryParticipateNumberResponse)
	},
}

// GetAlitripMerchantGalaxyQueryParticipateNumberResponse() 从对象池中获取AlitripMerchantGalaxyQueryParticipateNumberResponse
func GetAlitripMerchantGalaxyQueryParticipateNumberResponse() *AlitripMerchantGalaxyQueryParticipateNumberResponse {
	return poolAlitripMerchantGalaxyQueryParticipateNumberResponse.Get().(*AlitripMerchantGalaxyQueryParticipateNumberResponse)
}

// ReleaseAlitripMerchantGalaxyQueryParticipateNumberResponse 释放AlitripMerchantGalaxyQueryParticipateNumberResponse
func ReleaseAlitripMerchantGalaxyQueryParticipateNumberResponse(v *AlitripMerchantGalaxyQueryParticipateNumberResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyQueryParticipateNumberResponse.Put(v)
}
