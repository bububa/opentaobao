package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyHotelDetailSearchDataResponse 结构体
type AlitripMerchantGalaxyHotelDetailSearchDataResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误编码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	Content *HotelDetailInfoVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyHotelDetailSearchDataResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyHotelDetailSearchDataResponse)
	},
}

// GetAlitripMerchantGalaxyHotelDetailSearchDataResponse() 从对象池中获取AlitripMerchantGalaxyHotelDetailSearchDataResponse
func GetAlitripMerchantGalaxyHotelDetailSearchDataResponse() *AlitripMerchantGalaxyHotelDetailSearchDataResponse {
	return poolAlitripMerchantGalaxyHotelDetailSearchDataResponse.Get().(*AlitripMerchantGalaxyHotelDetailSearchDataResponse)
}

// ReleaseAlitripMerchantGalaxyHotelDetailSearchDataResponse 释放AlitripMerchantGalaxyHotelDetailSearchDataResponse
func ReleaseAlitripMerchantGalaxyHotelDetailSearchDataResponse(v *AlitripMerchantGalaxyHotelDetailSearchDataResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyHotelDetailSearchDataResponse.Put(v)
}
