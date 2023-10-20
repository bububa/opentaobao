package axindata

import (
	"sync"
)

// TaobaoAlitripTravelFscRouteApiProductAddTopResult 结构体
type TaobaoAlitripTravelFscRouteApiProductAddTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscRouteProductAddResponse *FscRouteProductAddResponse `json:"fsc_route_product_add_response,omitempty" xml:"fsc_route_product_add_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelFscRouteApiProductAddTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelFscRouteApiProductAddTopResult)
	},
}

// GetTaobaoAlitripTravelFscRouteApiProductAddTopResult() 从对象池中获取TaobaoAlitripTravelFscRouteApiProductAddTopResult
func GetTaobaoAlitripTravelFscRouteApiProductAddTopResult() *TaobaoAlitripTravelFscRouteApiProductAddTopResult {
	return poolTaobaoAlitripTravelFscRouteApiProductAddTopResult.Get().(*TaobaoAlitripTravelFscRouteApiProductAddTopResult)
}

// ReleaseTaobaoAlitripTravelFscRouteApiProductAddTopResult 释放TaobaoAlitripTravelFscRouteApiProductAddTopResult
func ReleaseTaobaoAlitripTravelFscRouteApiProductAddTopResult(v *TaobaoAlitripTravelFscRouteApiProductAddTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.FscRouteProductAddResponse = nil
	v.Success = false
	poolTaobaoAlitripTravelFscRouteApiProductAddTopResult.Put(v)
}
