package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 兑换结果
	Content *DerbyVoucherCardRedeemVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardRedeemResponse.Put(v)
}
