package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyCouponInvalidListResponse 结构体
type AlitripMerchantGalaxyCouponInvalidListResponse struct {
	// 优惠券详情
	Contents []CouponInfo `json:"contents,omitempty" xml:"contents>coupon_info,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyCouponInvalidListResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCouponInvalidListResponse)
	},
}

// GetAlitripMerchantGalaxyCouponInvalidListResponse() 从对象池中获取AlitripMerchantGalaxyCouponInvalidListResponse
func GetAlitripMerchantGalaxyCouponInvalidListResponse() *AlitripMerchantGalaxyCouponInvalidListResponse {
	return poolAlitripMerchantGalaxyCouponInvalidListResponse.Get().(*AlitripMerchantGalaxyCouponInvalidListResponse)
}

// ReleaseAlitripMerchantGalaxyCouponInvalidListResponse 释放AlitripMerchantGalaxyCouponInvalidListResponse
func ReleaseAlitripMerchantGalaxyCouponInvalidListResponse(v *AlitripMerchantGalaxyCouponInvalidListResponse) {
	v.Contents = v.Contents[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyCouponInvalidListResponse.Put(v)
}
