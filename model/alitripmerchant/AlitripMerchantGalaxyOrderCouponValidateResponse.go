package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOrderCouponValidateResponse 结构体
type AlitripMerchantGalaxyOrderCouponValidateResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 试单结果
	Content *CouponValidateOrderDto `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyOrderCouponValidateResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderCouponValidateResponse)
	},
}

// GetAlitripMerchantGalaxyOrderCouponValidateResponse() 从对象池中获取AlitripMerchantGalaxyOrderCouponValidateResponse
func GetAlitripMerchantGalaxyOrderCouponValidateResponse() *AlitripMerchantGalaxyOrderCouponValidateResponse {
	return poolAlitripMerchantGalaxyOrderCouponValidateResponse.Get().(*AlitripMerchantGalaxyOrderCouponValidateResponse)
}

// ReleaseAlitripMerchantGalaxyOrderCouponValidateResponse 释放AlitripMerchantGalaxyOrderCouponValidateResponse
func ReleaseAlitripMerchantGalaxyOrderCouponValidateResponse(v *AlitripMerchantGalaxyOrderCouponValidateResponse) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyOrderCouponValidateResponse.Put(v)
}
