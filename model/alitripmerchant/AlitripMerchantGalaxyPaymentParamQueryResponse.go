package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyPaymentParamQueryResponse 结构体
type AlitripMerchantGalaxyPaymentParamQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 支付参数
	Content *OrderPayDto `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyPaymentParamQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyPaymentParamQueryResponse)
	},
}

// GetAlitripMerchantGalaxyPaymentParamQueryResponse() 从对象池中获取AlitripMerchantGalaxyPaymentParamQueryResponse
func GetAlitripMerchantGalaxyPaymentParamQueryResponse() *AlitripMerchantGalaxyPaymentParamQueryResponse {
	return poolAlitripMerchantGalaxyPaymentParamQueryResponse.Get().(*AlitripMerchantGalaxyPaymentParamQueryResponse)
}

// ReleaseAlitripMerchantGalaxyPaymentParamQueryResponse 释放AlitripMerchantGalaxyPaymentParamQueryResponse
func ReleaseAlitripMerchantGalaxyPaymentParamQueryResponse(v *AlitripMerchantGalaxyPaymentParamQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyPaymentParamQueryResponse.Put(v)
}
