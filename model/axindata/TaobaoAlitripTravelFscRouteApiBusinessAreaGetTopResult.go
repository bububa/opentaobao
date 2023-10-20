package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult 结构体
type TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscBusinessAreaApiResponse *FscBusinessAreaApiResponse `json:"fsc_business_area_api_response,omitempty" xml:"fsc_business_area_api_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult
func GetTaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult() *TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult {
	return poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult.Get().(*TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult 释放TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult(v *TaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.FscBusinessAreaApiResponse = nil
	v.Success = false
	poolTaobaoAlitripTravelFscRouteApiBusinessAreaGetTopResult.Put(v)
}
