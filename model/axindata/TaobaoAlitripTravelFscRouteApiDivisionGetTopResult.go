package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiDivisionGetTopResult 结构体
type TaobaoAlitripTravelFscRouteApiDivisionGetTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscTripDivisionApiResponse *FscTripDivisionApiResponse `json:"fsc_trip_division_api_response,omitempty" xml:"fsc_trip_division_api_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiDivisionGetTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiDivisionGetTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiDivisionGetTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiDivisionGetTopResult
func GetTaobaoAlitripTravelFscRouteApiDivisionGetTopResult() *TaobaoAlitripTravelFscRouteApiDivisionGetTopResult {
	return poolTaobaoAlitripTravelFscRouteApiDivisionGetTopResult.Get().(*TaobaoAlitripTravelFscRouteApiDivisionGetTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiDivisionGetTopResult 释放TaobaoAlitripTravelFscRouteApiDivisionGetTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiDivisionGetTopResult(v *TaobaoAlitripTravelFscRouteApiDivisionGetTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.FscTripDivisionApiResponse = nil
	v.Success = false
	poolTaobaoAlitripTravelFscRouteApiDivisionGetTopResult.Put(v)
}
