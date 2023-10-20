package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProductUpdateTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProductUpdateTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProductUpdateTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductUpdateTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductUpdateTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProductUpdateTopResult
func GetTaobaoAlitripTravelFscRouteApiProductUpdateTopResult() *TaobaoAlitripTravelFscRouteApiProductUpdateTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProductUpdateTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProductUpdateTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductUpdateTopResult 释放TaobaoAlitripTravelFscRouteApiProductUpdateTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProductUpdateTopResult(v *TaobaoAlitripTravelFscRouteApiProductUpdateTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiProductUpdateTopResult.Put(v)
}
