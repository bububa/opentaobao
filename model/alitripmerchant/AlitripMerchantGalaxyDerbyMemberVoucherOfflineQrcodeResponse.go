package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse struct {
	// 200/500
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 权益券实体类
	Content *DerbyVoucherVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功/失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse() 从对象池中获取AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse() *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse 释放AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse.Put(v)
}
