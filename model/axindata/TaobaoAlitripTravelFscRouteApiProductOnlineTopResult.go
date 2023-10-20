package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProductOnlineTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProductOnlineTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProductOnlineTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductOnlineTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductOnlineTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProductOnlineTopResult
func GetTaobaoAlitripTravelFscRouteApiProductOnlineTopResult() *TaobaoAlitripTravelFscRouteApiProductOnlineTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProductOnlineTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProductOnlineTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductOnlineTopResult 释放TaobaoAlitripTravelFscRouteApiProductOnlineTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProductOnlineTopResult(v *TaobaoAlitripTravelFscRouteApiProductOnlineTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiProductOnlineTopResult.Put(v)
}
