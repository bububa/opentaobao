package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyVoucherQueryListResponse 结构体
type AlitripMerchantGalaxyVoucherQueryListResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回体
	Content *VoucherParameter `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyVoucherQueryListResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyVoucherQueryListResponse)
	},
}

// GetAlitripMerchantGalaxyVoucherQueryListResponse() 从对象池中获取AlitripMerchantGalaxyVoucherQueryListResponse
func GetAlitripMerchantGalaxyVoucherQueryListResponse() *AlitripMerchantGalaxyVoucherQueryListResponse {
	return poolAlitripMerchantGalaxyVoucherQueryListResponse.Get().(*AlitripMerchantGalaxyVoucherQueryListResponse)
}

// ReleaseAlitripMerchantGalaxyVoucherQueryListResponse 释放AlitripMerchantGalaxyVoucherQueryListResponse
func ReleaseAlitripMerchantGalaxyVoucherQueryListResponse(v *AlitripMerchantGalaxyVoucherQueryListResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyVoucherQueryListResponse.Put(v)
}
