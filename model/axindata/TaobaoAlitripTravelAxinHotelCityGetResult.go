package axindata

import (
	"sync"
)

// TaobaoAlitripTravelAxinHotelCityGetResult 结构体
type TaobaoAlitripTravelAxinHotelCityGetResult struct {
	// 城市信息
	DataList []HotelCityVo `json:"data_list,omitempty" xml:"data_list>hotel_city_vo,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelAxinHotelCityGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelCityGetResult)
	},
}

// GetTaobaoAlitripTravelAxinHotelCityGetResult() 从对象池中获取TaobaoAlitripTravelAxinHotelCityGetResult
func GetTaobaoAlitripTravelAxinHotelCityGetResult() *TaobaoAlitripTravelAxinHotelCityGetResult {
	return poolTaobaoAlitripTravelAxinHotelCityGetResult.Get().(*TaobaoAlitripTravelAxinHotelCityGetResult)
}

// ReleaseTaobaoAlitripTravelAxinHotelCityGetResult 释放TaobaoAlitripTravelAxinHotelCityGetResult
func ReleaseTaobaoAlitripTravelAxinHotelCityGetResult(v *TaobaoAlitripTravelAxinHotelCityGetResult) {
	v.DataList = v.DataList[:0]
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolTaobaoAlitripTravelAxinHotelCityGetResult.Put(v)
}
