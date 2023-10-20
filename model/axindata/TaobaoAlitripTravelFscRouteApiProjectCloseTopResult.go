package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProjectCloseTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProjectCloseTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProjectCloseTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectCloseTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectCloseTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProjectCloseTopResult
func GetTaobaoAlitripTravelFscRouteApiProjectCloseTopResult() *TaobaoAlitripTravelFscRouteApiProjectCloseTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProjectCloseTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProjectCloseTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectCloseTopResult 释放TaobaoAlitripTravelFscRouteApiProjectCloseTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProjectCloseTopResult(v *TaobaoAlitripTravelFscRouteApiProjectCloseTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiProjectCloseTopResult.Put(v)
}
