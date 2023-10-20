package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 实体类
	Content *DerbyMemberVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功/失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse() *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse.Put(v)
}
