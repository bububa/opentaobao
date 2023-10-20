package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyWechatCardParmQueryResponse 结构体
type AlitripMerchantGalaxyWechatCardParmQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 添加会员参数对象
	Content *MemberCardParamVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyWechatCardParmQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatCardParmQueryResponse)
	},
}

// GetAlitripMerchantGalaxyWechatCardParmQueryResponse() 从对象池中获取AlitripMerchantGalaxyWechatCardParmQueryResponse
func GetAlitripMerchantGalaxyWechatCardParmQueryResponse() *AlitripMerchantGalaxyWechatCardParmQueryResponse {
	return poolAlitripMerchantGalaxyWechatCardParmQueryResponse.Get().(*AlitripMerchantGalaxyWechatCardParmQueryResponse)
}

// ReleaseAlitripMerchantGalaxyWechatCardParmQueryResponse 释放AlitripMerchantGalaxyWechatCardParmQueryResponse
func ReleaseAlitripMerchantGalaxyWechatCardParmQueryResponse(v *AlitripMerchantGalaxyWechatCardParmQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyWechatCardParmQueryResponse.Put(v)
}
