package normalvisa

import (
	"sync"
)

// TaobaoAlitripTravelNormalvisaStoreuserResultSet 结构体
type TaobaoAlitripTravelNormalvisaStoreuserResultSet struct {
	// 结果：数字数组
	Results []int64 `json:"results,omitempty" xml:"results>int64,omitempty"`
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

var poolTaobaoAlitripTravelNormalvisaStoreuserResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelNormalvisaStoreuserResultSet)
	},
}

// GetTaobaoAlitripTravelNormalvisaStoreuserResultSet() 从对象池中获取TaobaoAlitripTravelNormalvisaStoreuserResultSet
func GetTaobaoAlitripTravelNormalvisaStoreuserResultSet() *TaobaoAlitripTravelNormalvisaStoreuserResultSet {
	return poolTaobaoAlitripTravelNormalvisaStoreuserResultSet.Get().(*TaobaoAlitripTravelNormalvisaStoreuserResultSet)
}

// ReleaseTaobaoAlitripTravelNormalvisaStoreuserResultSet 释放TaobaoAlitripTravelNormalvisaStoreuserResultSet
func ReleaseTaobaoAlitripTravelNormalvisaStoreuserResultSet(v *TaobaoAlitripTravelNormalvisaStoreuserResultSet) {
	v.Results = v.Results[:0]
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	v.HasNext = false
	poolTaobaoAlitripTravelNormalvisaStoreuserResultSet.Put(v)
}
