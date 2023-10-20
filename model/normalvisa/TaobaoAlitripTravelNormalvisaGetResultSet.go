package normalvisa

import (
	"sync"
)

// TaobaoAlitripTravelNormalvisaGetResultSet 结构体
type TaobaoAlitripTravelNormalvisaGetResultSet struct {
	// 结果
	Results []NormalVisaInfo `json:"results,omitempty" xml:"results>normal_visa_info,omitempty"`
	// 结果数目
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

var poolTaobaoAlitripTravelNormalvisaGetResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaGetResultSet)
	},
}

// GetTaobaoAlitripTravelNormalvisaGetResultSet() 从对象池中获取TaobaoAlitripTravelNormalvisaGetResultSet
func GetTaobaoAlitripTravelNormalvisaGetResultSet() *TaobaoAlitripTravelNormalvisaGetResultSet {
	return poolTaobaoAlitripTravelNormalvisaGetResultSet.Get().(*TaobaoAlitripTravelNormalvisaGetResultSet)
}

// ReleaseTaobaoAlitripTravelNormalvisaGetResultSet 释放TaobaoAlitripTravelNormalvisaGetResultSet
func ReleaseTaobaoAlitripTravelNormalvisaGetResultSet(v *TaobaoAlitripTravelNormalvisaGetResultSet) {
	v.Results = v.Results[:0]
	v.TotalResults = 0
	poolTaobaoAlitripTravelNormalvisaGetResultSet.Put(v)
}
