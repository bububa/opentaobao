package axintrade

import (
	"sync"
)

// TaobaoAlitripTravelAxinHotelOrderDetailResult 结构体
type TaobaoAlitripTravelAxinHotelOrderDetailResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回参数
	Data *HotelOrderQueryRes `json:"data,omitempty" xml:"data,omitempty"`
	// 成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelAxinHotelOrderDetailResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinHotelOrderDetailResult)
	},
}

// GetTaobaoAlitripTravelAxinHotelOrderDetailResult() 从对象池中获取TaobaoAlitripTravelAxinHotelOrderDetailResult
func GetTaobaoAlitripTravelAxinHotelOrderDetailResult() *TaobaoAlitripTravelAxinHotelOrderDetailResult {
	return poolTaobaoAlitripTravelAxinHotelOrderDetailResult.Get().(*TaobaoAlitripTravelAxinHotelOrderDetailResult)
}

// ReleaseTaobaoAlitripTravelAxinHotelOrderDetailResult 释放TaobaoAlitripTravelAxinHotelOrderDetailResult
func ReleaseTaobaoAlitripTravelAxinHotelOrderDetailResult(v *TaobaoAlitripTravelAxinHotelOrderDetailResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolTaobaoAlitripTravelAxinHotelOrderDetailResult.Put(v)
}
