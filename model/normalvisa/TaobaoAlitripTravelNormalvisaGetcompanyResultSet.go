package normalvisa

import (
	"sync"
)

// TaobaoAlitripTravelNormalvisaGetcompanyResultSet 结构体
type TaobaoAlitripTravelNormalvisaGetcompanyResultSet struct {
	// 结果
	Results []LogisticsCompanyInfo `json:"results,omitempty" xml:"results>logistics_company_info,omitempty"`
	// 异常
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否包含下一个
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolTaobaoAlitripTravelNormalvisaGetcompanyResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaGetcompanyResultSet)
	},
}

// GetTaobaoAlitripTravelNormalvisaGetcompanyResultSet() 从对象池中获取TaobaoAlitripTravelNormalvisaGetcompanyResultSet
func GetTaobaoAlitripTravelNormalvisaGetcompanyResultSet() *TaobaoAlitripTravelNormalvisaGetcompanyResultSet {
	return poolTaobaoAlitripTravelNormalvisaGetcompanyResultSet.Get().(*TaobaoAlitripTravelNormalvisaGetcompanyResultSet)
}

// ReleaseTaobaoAlitripTravelNormalvisaGetcompanyResultSet 释放TaobaoAlitripTravelNormalvisaGetcompanyResultSet
func ReleaseTaobaoAlitripTravelNormalvisaGetcompanyResultSet(v *TaobaoAlitripTravelNormalvisaGetcompanyResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	v.HasNext = false
	poolTaobaoAlitripTravelNormalvisaGetcompanyResultSet.Put(v)
}
