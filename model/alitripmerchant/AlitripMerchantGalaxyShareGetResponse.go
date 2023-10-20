package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyShareGetResponse 结构体
type AlitripMerchantGalaxyShareGetResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回分享素材
	Content *ShareMaterialVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyShareGetResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyShareGetResponse)
	},
}

// GetAlitripMerchantGalaxyShareGetResponse() 从对象池中获取AlitripMerchantGalaxyShareGetResponse
func GetAlitripMerchantGalaxyShareGetResponse() *AlitripMerchantGalaxyShareGetResponse {
	return poolAlitripMerchantGalaxyShareGetResponse.Get().(*AlitripMerchantGalaxyShareGetResponse)
}

// ReleaseAlitripMerchantGalaxyShareGetResponse 释放AlitripMerchantGalaxyShareGetResponse
func ReleaseAlitripMerchantGalaxyShareGetResponse(v *AlitripMerchantGalaxyShareGetResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyShareGetResponse.Put(v)
}
