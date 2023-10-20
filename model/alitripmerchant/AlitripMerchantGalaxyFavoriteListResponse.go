package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyFavoriteListResponse 结构体
type AlitripMerchantGalaxyFavoriteListResponse struct {
	// 收藏列表
	Contents []FavoriteHotelList `json:"contents,omitempty" xml:"contents>favorite_hotel_list,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyFavoriteListResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyFavoriteListResponse)
	},
}

// GetAlitripMerchantGalaxyFavoriteListResponse() 从对象池中获取AlitripMerchantGalaxyFavoriteListResponse
func GetAlitripMerchantGalaxyFavoriteListResponse() *AlitripMerchantGalaxyFavoriteListResponse {
	return poolAlitripMerchantGalaxyFavoriteListResponse.Get().(*AlitripMerchantGalaxyFavoriteListResponse)
}

// ReleaseAlitripMerchantGalaxyFavoriteListResponse 释放AlitripMerchantGalaxyFavoriteListResponse
func ReleaseAlitripMerchantGalaxyFavoriteListResponse(v *AlitripMerchantGalaxyFavoriteListResponse) {
	v.Contents = v.Contents[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripMerchantGalaxyFavoriteListResponse.Put(v)
}
