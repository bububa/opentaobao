package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyCouponValidListResponse 结构体
type AlitripMerchantGalaxyCouponValidListResponse struct {
	// 状态码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 优惠券实例
	Content *WeChatCouponVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyCouponValidListResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCouponValidListResponse)
	},
}

// GetAlitripMerchantGalaxyCouponValidListResponse() 从对象池中获取AlitripMerchantGalaxyCouponValidListResponse
func GetAlitripMerchantGalaxyCouponValidListResponse() *AlitripMerchantGalaxyCouponValidListResponse {
	return poolAlitripMerchantGalaxyCouponValidListResponse.Get().(*AlitripMerchantGalaxyCouponValidListResponse)
}

// ReleaseAlitripMerchantGalaxyCouponValidListResponse 释放AlitripMerchantGalaxyCouponValidListResponse
func ReleaseAlitripMerchantGalaxyCouponValidListResponse(v *AlitripMerchantGalaxyCouponValidListResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyCouponValidListResponse.Put(v)
}
