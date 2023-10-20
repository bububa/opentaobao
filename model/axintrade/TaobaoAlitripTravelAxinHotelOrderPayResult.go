package axintrade

import (
	"sync"
)

// TaobaoAlitripTravelAxinHotelOrderPayResult 结构体
type TaobaoAlitripTravelAxinHotelOrderPayResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 支付id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelAxinHotelOrderPayResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelOrderPayResult)
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderPayResult() 从对象池中获取TaobaoAlitripTravelAxinHotelOrderPayResult
func GetTaobaoAlitripTravelAxinHotelOrderPayResult() *TaobaoAlitripTravelAxinHotelOrderPayResult {
	return poolTaobaoAlitripTravelAxinHotelOrderPayResult.Get().(*TaobaoAlitripTravelAxinHotelOrderPayResult)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderPayResult 释放TaobaoAlitripTravelAxinHotelOrderPayResult
func ReleaseTaobaoAlitripTravelAxinHotelOrderPayResult(v *TaobaoAlitripTravelAxinHotelOrderPayResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = 0
	v.Success = false
	poolTaobaoAlitripTravelAxinHotelOrderPayResult.Put(v)
}
