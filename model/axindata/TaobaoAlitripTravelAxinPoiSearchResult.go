package axindata

import (
	"sync"
)

// TaobaoAlitripTravelAxinPoiSearchResult 结构体
type TaobaoAlitripTravelAxinPoiSearchResult struct {
	// 列表
	DataList []PoiVo `json:"data_list,omitempty" xml:"data_list>poi_vo,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripTravelAxinPoiSearchResult = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelAxinPoiSearchResult)
	},
}

// GetTaobaoAlitripTravelAxinPoiSearchResult() 从对象池中获取TaobaoAlitripTravelAxinPoiSearchResult
func GetTaobaoAlitripTravelAxinPoiSearchResult() *TaobaoAlitripTravelAxinPoiSearchResult {
	return poolTaobaoAlitripTravelAxinPoiSearchResult.Get().(*TaobaoAlitripTravelAxinPoiSearchResult)
}

// ReleaseTaobaoAlitripTravelAxinPoiSearchResult 释放TaobaoAlitripTravelAxinPoiSearchResult
func ReleaseTaobaoAlitripTravelAxinPoiSearchResult(v *TaobaoAlitripTravelAxinPoiSearchResult) {
	v.DataList = v.DataList[:0]
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolTaobaoAlitripTravelAxinPoiSearchResult.Put(v)
}
