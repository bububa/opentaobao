package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyCityListResponse 结构体
type AlitripMerchantGalaxyCityListResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回城市列表
	Content *AddressListSearchDto `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyCityListResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCityListResponse)
	},
}

// GetAlitripMerchantGalaxyCityListResponse() 从对象池中获取AlitripMerchantGalaxyCityListResponse
func GetAlitripMerchantGalaxyCityListResponse() *AlitripMerchantGalaxyCityListResponse {
	return poolAlitripMerchantGalaxyCityListResponse.Get().(*AlitripMerchantGalaxyCityListResponse)
}

// ReleaseAlitripMerchantGalaxyCityListResponse 释放AlitripMerchantGalaxyCityListResponse
func ReleaseAlitripMerchantGalaxyCityListResponse(v *AlitripMerchantGalaxyCityListResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyCityListResponse.Put(v)
}
