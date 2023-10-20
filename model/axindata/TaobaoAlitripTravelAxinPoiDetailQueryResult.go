package axindata

import (
	"sync"
)

// TaobaoAlitripTravelAxinPoiDetailQueryResult 结构体
type TaobaoAlitripTravelAxinPoiDetailQueryResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// poi详情
	Data *PoiDetaiVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelAxinPoiDetailQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinPoiDetailQueryResult)
	},
}

// GetTaobaoAlitripTravelAxinPoiDetailQueryResult() 从对象池中获取TaobaoAlitripTravelAxinPoiDetailQueryResult
func GetTaobaoAlitripTravelAxinPoiDetailQueryResult() *TaobaoAlitripTravelAxinPoiDetailQueryResult {
	return poolTaobaoAlitripTravelAxinPoiDetailQueryResult.Get().(*TaobaoAlitripTravelAxinPoiDetailQueryResult)
}

// ReleaseTaobaoAlitripTravelAxinPoiDetailQueryResult 释放TaobaoAlitripTravelAxinPoiDetailQueryResult
func ReleaseTaobaoAlitripTravelAxinPoiDetailQueryResult(v *TaobaoAlitripTravelAxinPoiDetailQueryResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolTaobaoAlitripTravelAxinPoiDetailQueryResult.Put(v)
}
