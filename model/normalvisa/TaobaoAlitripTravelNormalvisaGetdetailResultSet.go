package normalvisa

import (
	"sync"
)

// TaobaoAlitripTravelNormalvisaGetdetailResultSet 结构体
type TaobaoAlitripTravelNormalvisaGetdetailResultSet struct {
	// 结果
	Results []NormalVisaDetailInfo `json:"results,omitempty" xml:"results>normal_visa_detail_info,omitempty"`
	// 异常
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否包含下一个
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolTaobaoAlitripTravelNormalvisaGetdetailResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaGetdetailResultSet)
	},
}

// GetTaobaoAlitripTravelNormalvisaGetdetailResultSet() 从对象池中获取TaobaoAlitripTravelNormalvisaGetdetailResultSet
func GetTaobaoAlitripTravelNormalvisaGetdetailResultSet() *TaobaoAlitripTravelNormalvisaGetdetailResultSet {
	return poolTaobaoAlitripTravelNormalvisaGetdetailResultSet.Get().(*TaobaoAlitripTravelNormalvisaGetdetailResultSet)
}

// ReleaseTaobaoAlitripTravelNormalvisaGetdetailResultSet 释放TaobaoAlitripTravelNormalvisaGetdetailResultSet
func ReleaseTaobaoAlitripTravelNormalvisaGetdetailResultSet(v *TaobaoAlitripTravelNormalvisaGetdetailResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	v.HasNext = false
	poolTaobaoAlitripTravelNormalvisaGetdetailResultSet.Put(v)
}
