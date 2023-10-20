package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyCommonGetEnumsbynameResponse 结构体
type AlitripMerchantGalaxyCommonGetEnumsbynameResponse struct {
	// 枚举
	Content []EnumVo `json:"content,omitempty" xml:"content>enum_vo,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyCommonGetEnumsbynameResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCommonGetEnumsbynameResponse)
	},
}

// GetAlitripMerchantGalaxyCommonGetEnumsbynameResponse() 从对象池中获取AlitripMerchantGalaxyCommonGetEnumsbynameResponse
func GetAlitripMerchantGalaxyCommonGetEnumsbynameResponse() *AlitripMerchantGalaxyCommonGetEnumsbynameResponse {
	return poolAlitripMerchantGalaxyCommonGetEnumsbynameResponse.Get().(*AlitripMerchantGalaxyCommonGetEnumsbynameResponse)
}

// ReleaseAlitripMerchantGalaxyCommonGetEnumsbynameResponse 释放AlitripMerchantGalaxyCommonGetEnumsbynameResponse
func ReleaseAlitripMerchantGalaxyCommonGetEnumsbynameResponse(v *AlitripMerchantGalaxyCommonGetEnumsbynameResponse) {
	v.Content = v.Content[:0]
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlitripMerchantGalaxyCommonGetEnumsbynameResponse.Put(v)
}
