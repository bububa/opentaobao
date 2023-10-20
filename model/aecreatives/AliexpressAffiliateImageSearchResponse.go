package aecreatives

import (
	"sync"
)

// AliexpressAffiliateImageSearchResponse 结构体
type AliexpressAffiliateImageSearchResponse struct {
	// 返回结果状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 默认描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 图搜结果
	Data *TrafficImageSearchResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAliexpressAffiliateImageSearchResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateImageSearchResponse)
	},
}

// GetAliexpressAffiliateImageSearchResponse() 从对象池中获取AliexpressAffiliateImageSearchResponse
func GetAliexpressAffiliateImageSearchResponse() *AliexpressAffiliateImageSearchResponse {
	return poolAliexpressAffiliateImageSearchResponse.Get().(*AliexpressAffiliateImageSearchResponse)
}

// ReleaseAliexpressAffiliateImageSearchResponse 释放AliexpressAffiliateImageSearchResponse
func ReleaseAliexpressAffiliateImageSearchResponse(v *AliexpressAffiliateImageSearchResponse) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolAliexpressAffiliateImageSearchResponse.Put(v)
}
