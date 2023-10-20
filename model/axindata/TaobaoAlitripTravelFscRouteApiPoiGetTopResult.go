package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiPoiGetTopResult 结构体
type TaobaoAlitripTravelFscRouteApiPoiGetTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscPoiApiResponse *FscPoiApiResponse `json:"fsc_poi_api_response,omitempty" xml:"fsc_poi_api_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiPoiGetTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiPoiGetTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiPoiGetTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiPoiGetTopResult
func GetTaobaoAlitripTravelFscRouteApiPoiGetTopResult() *TaobaoAlitripTravelFscRouteApiPoiGetTopResult {
	return poolTaobaoAlitripTravelFscRouteApiPoiGetTopResult.Get().(*TaobaoAlitripTravelFscRouteApiPoiGetTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiPoiGetTopResult 释放TaobaoAlitripTravelFscRouteApiPoiGetTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiPoiGetTopResult(v *TaobaoAlitripTravelFscRouteApiPoiGetTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.FscPoiApiResponse = nil
	v.Success = false
	poolTaobaoAlitripTravelFscRouteApiPoiGetTopResult.Put(v)
}
