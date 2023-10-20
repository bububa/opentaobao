package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyVoucherQueryResponse 结构体
type AlitripMerchantGalaxyVoucherQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回信息
	Content *VoucherInfoVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyVoucherQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyVoucherQueryResponse)
	},
}

// GetAlitripMerchantGalaxyVoucherQueryResponse() 从对象池中获取AlitripMerchantGalaxyVoucherQueryResponse
func GetAlitripMerchantGalaxyVoucherQueryResponse() *AlitripMerchantGalaxyVoucherQueryResponse {
	return poolAlitripMerchantGalaxyVoucherQueryResponse.Get().(*AlitripMerchantGalaxyVoucherQueryResponse)
}

// ReleaseAlitripMerchantGalaxyVoucherQueryResponse 释放AlitripMerchantGalaxyVoucherQueryResponse
func ReleaseAlitripMerchantGalaxyVoucherQueryResponse(v *AlitripMerchantGalaxyVoucherQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyVoucherQueryResponse.Put(v)
}
