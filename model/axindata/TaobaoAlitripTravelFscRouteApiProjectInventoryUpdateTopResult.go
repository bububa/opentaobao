package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult
func GetTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult() *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult 释放TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult(v *TaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiProjectInventoryUpdateTopResult.Put(v)
}
