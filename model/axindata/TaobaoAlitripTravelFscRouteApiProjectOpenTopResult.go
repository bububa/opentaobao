package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProjectOpenTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProjectOpenTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProjectOpenTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectOpenTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectOpenTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProjectOpenTopResult
func GetTaobaoAlitripTravelFscRouteApiProjectOpenTopResult() *TaobaoAlitripTravelFscRouteApiProjectOpenTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProjectOpenTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProjectOpenTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectOpenTopResult 释放TaobaoAlitripTravelFscRouteApiProjectOpenTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProjectOpenTopResult(v *TaobaoAlitripTravelFscRouteApiProjectOpenTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiProjectOpenTopResult.Put(v)
}
