package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyUserRiskResponse 结构体
type AlitripMerchantGalaxyUserRiskResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回内容
	Content *UserRiskRankVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyUserRiskResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyUserRiskResponse)
	},
}

// GetAlitripMerchantGalaxyUserRiskResponse() 从对象池中获取AlitripMerchantGalaxyUserRiskResponse
func GetAlitripMerchantGalaxyUserRiskResponse() *AlitripMerchantGalaxyUserRiskResponse {
	return poolAlitripMerchantGalaxyUserRiskResponse.Get().(*AlitripMerchantGalaxyUserRiskResponse)
}

// ReleaseAlitripMerchantGalaxyUserRiskResponse 释放AlitripMerchantGalaxyUserRiskResponse
func ReleaseAlitripMerchantGalaxyUserRiskResponse(v *AlitripMerchantGalaxyUserRiskResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyUserRiskResponse.Put(v)
}
