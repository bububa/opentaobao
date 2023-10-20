package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyActivityCouponListResponse 结构体
type AlitripMerchantGalaxyActivityCouponListResponse struct {
	// 活动详情
	Contents []CouponActivity `json:"contents,omitempty" xml:"contents>coupon_activity,omitempty"`
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyActivityCouponListResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityCouponListResponse)
	},
}

// GetAlitripMerchantGalaxyActivityCouponListResponse() 从对象池中获取AlitripMerchantGalaxyActivityCouponListResponse
func GetAlitripMerchantGalaxyActivityCouponListResponse() *AlitripMerchantGalaxyActivityCouponListResponse {
	return poolAlitripMerchantGalaxyActivityCouponListResponse.Get().(*AlitripMerchantGalaxyActivityCouponListResponse)
}

// ReleaseAlitripMerchantGalaxyActivityCouponListResponse 释放AlitripMerchantGalaxyActivityCouponListResponse
func ReleaseAlitripMerchantGalaxyActivityCouponListResponse(v *AlitripMerchantGalaxyActivityCouponListResponse) {
	v.Contents = v.Contents[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyActivityCouponListResponse.Put(v)
}
