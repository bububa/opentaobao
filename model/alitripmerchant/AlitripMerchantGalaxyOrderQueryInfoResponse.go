package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOrderQueryInfoResponse 结构体
type AlitripMerchantGalaxyOrderQueryInfoResponse struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息展示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单详情页对象
	Content *VoucherOrderDetailVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyOrderQueryInfoResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderQueryInfoResponse)
	},
}

// GetAlitripMerchantGalaxyOrderQueryInfoResponse() 从对象池中获取AlitripMerchantGalaxyOrderQueryInfoResponse
func GetAlitripMerchantGalaxyOrderQueryInfoResponse() *AlitripMerchantGalaxyOrderQueryInfoResponse {
	return poolAlitripMerchantGalaxyOrderQueryInfoResponse.Get().(*AlitripMerchantGalaxyOrderQueryInfoResponse)
}

// ReleaseAlitripMerchantGalaxyOrderQueryInfoResponse 释放AlitripMerchantGalaxyOrderQueryInfoResponse
func ReleaseAlitripMerchantGalaxyOrderQueryInfoResponse(v *AlitripMerchantGalaxyOrderQueryInfoResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyOrderQueryInfoResponse.Put(v)
}
