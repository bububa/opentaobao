package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscProductLabelApiResponse *FscProductLabelApiResponse `json:"fsc_product_label_api_response,omitempty" xml:"fsc_product_label_api_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProductLabelGetTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductLabelGetTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult
func GetTaobaoAlitripTravelFscRouteApiProductLabelGetTopResult() *TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProductLabelGetTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductLabelGetTopResult 释放TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProductLabelGetTopResult(v *TaobaoAlitripTravelFscRouteApiProductLabelGetTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.FscProductLabelApiResponse = nil
	v.Success = false
	poolTaobaoAlitripTravelFscRouteApiProductLabelGetTopResult.Put(v)
}
