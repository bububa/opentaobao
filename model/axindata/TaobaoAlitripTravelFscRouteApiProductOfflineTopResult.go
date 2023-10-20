package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProductOfflineTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProductOfflineTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProductOfflineTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductOfflineTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductOfflineTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProductOfflineTopResult
func GetTaobaoAlitripTravelFscRouteApiProductOfflineTopResult() *TaobaoAlitripTravelFscRouteApiProductOfflineTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProductOfflineTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProductOfflineTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductOfflineTopResult 释放TaobaoAlitripTravelFscRouteApiProductOfflineTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProductOfflineTopResult(v *TaobaoAlitripTravelFscRouteApiProductOfflineTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Data = false
	poolTaobaoAlitripTravelFscRouteApiProductOfflineTopResult.Put(v)
}
