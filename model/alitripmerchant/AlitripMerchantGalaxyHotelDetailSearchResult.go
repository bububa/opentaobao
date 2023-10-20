package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyHotelDetailSearchResult 结构体
type AlitripMerchantGalaxyHotelDetailSearchResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 酒店详情返回实体
	Content *HotelDetailInfoDto `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyHotelDetailSearchResult = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyHotelDetailSearchResult)
	},
}

// GetAlitripMerchantGalaxyHotelDetailSearchResult() 从对象池中获取AlitripMerchantGalaxyHotelDetailSearchResult
func GetAlitripMerchantGalaxyHotelDetailSearchResult() *AlitripMerchantGalaxyHotelDetailSearchResult {
	return poolAlitripMerchantGalaxyHotelDetailSearchResult.Get().(*AlitripMerchantGalaxyHotelDetailSearchResult)
}

// ReleaseAlitripMerchantGalaxyHotelDetailSearchResult 释放AlitripMerchantGalaxyHotelDetailSearchResult
func ReleaseAlitripMerchantGalaxyHotelDetailSearchResult(v *AlitripMerchantGalaxyHotelDetailSearchResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyHotelDetailSearchResult.Put(v)
}
