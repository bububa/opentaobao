package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiPoiApplyTopResult 结构体
type TaobaoAlitripTravelFscRouteApiPoiApplyTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiPoiApplyTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiPoiApplyTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiPoiApplyTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiPoiApplyTopResult
func GetTaobaoAlitripTravelFscRouteApiPoiApplyTopResult() *TaobaoAlitripTravelFscRouteApiPoiApplyTopResult {
	return poolTaobaoAlitripTravelFscRouteApiPoiApplyTopResult.Get().(*TaobaoAlitripTravelFscRouteApiPoiApplyTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiPoiApplyTopResult 释放TaobaoAlitripTravelFscRouteApiPoiApplyTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiPoiApplyTopResult(v *TaobaoAlitripTravelFscRouteApiPoiApplyTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiPoiApplyTopResult.Put(v)
}
