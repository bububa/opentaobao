package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult 结构体
type TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiDivisionApplyTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiDivisionApplyTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult
func GetTaobaoAlitripTravelFscRouteApiDivisionApplyTopResult() *TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult {
	return poolTaobaoAlitripTravelFscRouteApiDivisionApplyTopResult.Get().(*TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiDivisionApplyTopResult 释放TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiDivisionApplyTopResult(v *TaobaoAlitripTravelFscRouteApiDivisionApplyTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiDivisionApplyTopResult.Put(v)
}
