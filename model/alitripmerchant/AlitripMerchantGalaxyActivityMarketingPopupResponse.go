package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyActivityMarketingPopupResponse 结构体
type AlitripMerchantGalaxyActivityMarketingPopupResponse struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	Content *PopUpInfoVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyActivityMarketingPopupResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityMarketingPopupResponse)
	},
}

// GetAlitripMerchantGalaxyActivityMarketingPopupResponse() 从对象池中获取AlitripMerchantGalaxyActivityMarketingPopupResponse
func GetAlitripMerchantGalaxyActivityMarketingPopupResponse() *AlitripMerchantGalaxyActivityMarketingPopupResponse {
	return poolAlitripMerchantGalaxyActivityMarketingPopupResponse.Get().(*AlitripMerchantGalaxyActivityMarketingPopupResponse)
}

// ReleaseAlitripMerchantGalaxyActivityMarketingPopupResponse 释放AlitripMerchantGalaxyActivityMarketingPopupResponse
func ReleaseAlitripMerchantGalaxyActivityMarketingPopupResponse(v *AlitripMerchantGalaxyActivityMarketingPopupResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyActivityMarketingPopupResponse.Put(v)
}
