package normalvisa

import (
	"sync"
)

// TaobaoAlitripTravelNormalvisaUploadfileResultSet 结构体
type TaobaoAlitripTravelNormalvisaUploadfileResultSet struct {
	// 结果列表
	Results []bool `json:"results,omitempty" xml:"results>bool,omitempty"`
	// 异常
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否包含下一个
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolTaobaoAlitripTravelNormalvisaUploadfileResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaUploadfileResultSet)
	},
}

// GetTaobaoAlitripTravelNormalvisaUploadfileResultSet() 从对象池中获取TaobaoAlitripTravelNormalvisaUploadfileResultSet
func GetTaobaoAlitripTravelNormalvisaUploadfileResultSet() *TaobaoAlitripTravelNormalvisaUploadfileResultSet {
	return poolTaobaoAlitripTravelNormalvisaUploadfileResultSet.Get().(*TaobaoAlitripTravelNormalvisaUploadfileResultSet)
}

// ReleaseTaobaoAlitripTravelNormalvisaUploadfileResultSet 释放TaobaoAlitripTravelNormalvisaUploadfileResultSet
func ReleaseTaobaoAlitripTravelNormalvisaUploadfileResultSet(v *TaobaoAlitripTravelNormalvisaUploadfileResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	v.HasNext = false
	poolTaobaoAlitripTravelNormalvisaUploadfileResultSet.Put(v)
}
