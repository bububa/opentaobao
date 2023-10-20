package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyFavoriteSaveResponse 结构体
type AlitripMerchantGalaxyFavoriteSaveResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 收藏状态
	Content *FavoriteStatusVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyFavoriteSaveResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyFavoriteSaveResponse)
	},
}

// GetAlitripMerchantGalaxyFavoriteSaveResponse() 从对象池中获取AlitripMerchantGalaxyFavoriteSaveResponse
func GetAlitripMerchantGalaxyFavoriteSaveResponse() *AlitripMerchantGalaxyFavoriteSaveResponse {
	return poolAlitripMerchantGalaxyFavoriteSaveResponse.Get().(*AlitripMerchantGalaxyFavoriteSaveResponse)
}

// ReleaseAlitripMerchantGalaxyFavoriteSaveResponse 释放AlitripMerchantGalaxyFavoriteSaveResponse
func ReleaseAlitripMerchantGalaxyFavoriteSaveResponse(v *AlitripMerchantGalaxyFavoriteSaveResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyFavoriteSaveResponse.Put(v)
}
