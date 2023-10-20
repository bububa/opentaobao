package axindata

import (
	"sync"
)

// TaobaoAlitripTravelAxinHotelPriceGetResult 结构体
type TaobaoAlitripTravelAxinHotelPriceGetResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回素材id
	Data *HotelPriceVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelAxinHotelPriceGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelPriceGetResult)
	},
}

// GetTaobaoAlitripTravelAxinHotelPriceGetResult() 从对象池中获取TaobaoAlitripTravelAxinHotelPriceGetResult
func GetTaobaoAlitripTravelAxinHotelPriceGetResult() *TaobaoAlitripTravelAxinHotelPriceGetResult {
	return poolTaobaoAlitripTravelAxinHotelPriceGetResult.Get().(*TaobaoAlitripTravelAxinHotelPriceGetResult)
}

// ReleaseTaobaoAlitripTravelAxinHotelPriceGetResult 释放TaobaoAlitripTravelAxinHotelPriceGetResult
func ReleaseTaobaoAlitripTravelAxinHotelPriceGetResult(v *TaobaoAlitripTravelAxinHotelPriceGetResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolTaobaoAlitripTravelAxinHotelPriceGetResult.Put(v)
}
