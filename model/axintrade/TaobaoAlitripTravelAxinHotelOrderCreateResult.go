package axintrade

import (
	"sync"
)

// TaobaoAlitripTravelAxinHotelOrderCreateResult 结构体
type TaobaoAlitripTravelAxinHotelOrderCreateResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 阿信订单id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelAxinHotelOrderCreateResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelOrderCreateResult)
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderCreateResult() 从对象池中获取TaobaoAlitripTravelAxinHotelOrderCreateResult
func GetTaobaoAlitripTravelAxinHotelOrderCreateResult() *TaobaoAlitripTravelAxinHotelOrderCreateResult {
	return poolTaobaoAlitripTravelAxinHotelOrderCreateResult.Get().(*TaobaoAlitripTravelAxinHotelOrderCreateResult)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderCreateResult 释放TaobaoAlitripTravelAxinHotelOrderCreateResult
func ReleaseTaobaoAlitripTravelAxinHotelOrderCreateResult(v *TaobaoAlitripTravelAxinHotelOrderCreateResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Data = 0
	v.Success = false
	poolTaobaoAlitripTravelAxinHotelOrderCreateResult.Put(v)
}
