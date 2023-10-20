package axindata

import (
	"sync"
)

// TaobaoAlitripTravelAxinHotelListGetResult 结构体
type TaobaoAlitripTravelAxinHotelListGetResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 分页对象
	Data *PageVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelAxinHotelListGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelListGetResult)
	},
}

// GetTaobaoAlitripTravelAxinHotelListGetResult() 从对象池中获取TaobaoAlitripTravelAxinHotelListGetResult
func GetTaobaoAlitripTravelAxinHotelListGetResult() *TaobaoAlitripTravelAxinHotelListGetResult {
	return poolTaobaoAlitripTravelAxinHotelListGetResult.Get().(*TaobaoAlitripTravelAxinHotelListGetResult)
}

// ReleaseTaobaoAlitripTravelAxinHotelListGetResult 释放TaobaoAlitripTravelAxinHotelListGetResult
func ReleaseTaobaoAlitripTravelAxinHotelListGetResult(v *TaobaoAlitripTravelAxinHotelListGetResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolTaobaoAlitripTravelAxinHotelListGetResult.Put(v)
}
