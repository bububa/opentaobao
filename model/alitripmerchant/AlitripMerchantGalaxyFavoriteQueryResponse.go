package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyFavoriteQueryResponse 结构体
type AlitripMerchantGalaxyFavoriteQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 收藏状态
	Content *FavoriteStatusVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyFavoriteQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyFavoriteQueryResponse)
	},
}

// GetAlitripMerchantGalaxyFavoriteQueryResponse() 从对象池中获取AlitripMerchantGalaxyFavoriteQueryResponse
func GetAlitripMerchantGalaxyFavoriteQueryResponse() *AlitripMerchantGalaxyFavoriteQueryResponse {
	return poolAlitripMerchantGalaxyFavoriteQueryResponse.Get().(*AlitripMerchantGalaxyFavoriteQueryResponse)
}

// ReleaseAlitripMerchantGalaxyFavoriteQueryResponse 释放AlitripMerchantGalaxyFavoriteQueryResponse
func ReleaseAlitripMerchantGalaxyFavoriteQueryResponse(v *AlitripMerchantGalaxyFavoriteQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyFavoriteQueryResponse.Put(v)
}
