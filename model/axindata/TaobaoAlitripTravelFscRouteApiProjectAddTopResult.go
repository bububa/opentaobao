package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProjectAddTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProjectAddTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscProjectAddApiResponse *FscProjectAddApiResponse `json:"fsc_project_add_api_response,omitempty" xml:"fsc_project_add_api_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProjectAddTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProjectAddTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProjectAddTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProjectAddTopResult
func GetTaobaoAlitripTravelFscRouteApiProjectAddTopResult() *TaobaoAlitripTravelFscRouteApiProjectAddTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProjectAddTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProjectAddTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProjectAddTopResult 释放TaobaoAlitripTravelFscRouteApiProjectAddTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProjectAddTopResult(v *TaobaoAlitripTravelFscRouteApiProjectAddTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.FscProjectAddApiResponse = nil
	v.Success = false
	poolTaobaoAlitripTravelFscRouteApiProjectAddTopResult.Put(v)
}
