package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProjectUpdateTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectUpdateTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult
func GetTaobaoAlitripTravelFscRouteApiProjectUpdateTopResult() *TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProjectUpdateTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectUpdateTopResult 释放TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProjectUpdateTopResult(v *TaobaoAlitripTravelFscRouteApiProjectUpdateTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiProjectUpdateTopResult.Put(v)
}
