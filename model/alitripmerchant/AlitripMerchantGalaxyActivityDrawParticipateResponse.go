package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyActivityDrawParticipateResponse 结构体
type AlitripMerchantGalaxyActivityDrawParticipateResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回类型
	Content *ActivityParticipateVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyActivityDrawParticipateResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityDrawParticipateResponse)
	},
}

// GetAlitripMerchantGalaxyActivityDrawParticipateResponse() 从对象池中获取AlitripMerchantGalaxyActivityDrawParticipateResponse
func GetAlitripMerchantGalaxyActivityDrawParticipateResponse() *AlitripMerchantGalaxyActivityDrawParticipateResponse {
	return poolAlitripMerchantGalaxyActivityDrawParticipateResponse.Get().(*AlitripMerchantGalaxyActivityDrawParticipateResponse)
}

// ReleaseAlitripMerchantGalaxyActivityDrawParticipateResponse 释放AlitripMerchantGalaxyActivityDrawParticipateResponse
func ReleaseAlitripMerchantGalaxyActivityDrawParticipateResponse(v *AlitripMerchantGalaxyActivityDrawParticipateResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyActivityDrawParticipateResponse.Put(v)
}
