package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyQueryDrawSummaryResponse 结构体
type AlitripMerchantGalaxyQueryDrawSummaryResponse struct {
	// 返回类型
	Contents []DrawOfferSummaryVo `json:"contents,omitempty" xml:"contents>draw_offer_summary_vo,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyQueryDrawSummaryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyQueryDrawSummaryResponse)
	},
}

// GetAlitripMerchantGalaxyQueryDrawSummaryResponse() 从对象池中获取AlitripMerchantGalaxyQueryDrawSummaryResponse
func GetAlitripMerchantGalaxyQueryDrawSummaryResponse() *AlitripMerchantGalaxyQueryDrawSummaryResponse {
	return poolAlitripMerchantGalaxyQueryDrawSummaryResponse.Get().(*AlitripMerchantGalaxyQueryDrawSummaryResponse)
}

// ReleaseAlitripMerchantGalaxyQueryDrawSummaryResponse 释放AlitripMerchantGalaxyQueryDrawSummaryResponse
func ReleaseAlitripMerchantGalaxyQueryDrawSummaryResponse(v *AlitripMerchantGalaxyQueryDrawSummaryResponse) {
	v.Contents = v.Contents[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyQueryDrawSummaryResponse.Put(v)
}
