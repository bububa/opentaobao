package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 内容
	Content *DerbyVoucherCardOrderDetailsVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryResponse.Put(v)
}
