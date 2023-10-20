package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyCityLikeResponse 结构体
type AlitripMerchantGalaxyCityLikeResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	Content *AddressSearchVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyCityLikeResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCityLikeResponse)
	},
}

// GetAlitripMerchantGalaxyCityLikeResponse() 从对象池中获取AlitripMerchantGalaxyCityLikeResponse
func GetAlitripMerchantGalaxyCityLikeResponse() *AlitripMerchantGalaxyCityLikeResponse {
	return poolAlitripMerchantGalaxyCityLikeResponse.Get().(*AlitripMerchantGalaxyCityLikeResponse)
}

// ReleaseAlitripMerchantGalaxyCityLikeResponse 释放AlitripMerchantGalaxyCityLikeResponse
func ReleaseAlitripMerchantGalaxyCityLikeResponse(v *AlitripMerchantGalaxyCityLikeResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyCityLikeResponse.Put(v)
}
