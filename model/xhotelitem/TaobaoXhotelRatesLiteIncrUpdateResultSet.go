package xhotelitem

import (
	"sync"
)

// TaobaoXhotelRatesLiteIncrUpdateResultSet 结构体
type TaobaoXhotelRatesLiteIncrUpdateResultSet struct {
	// 多个rate的更新结果
	FirstResult string `json:"first_result,omitempty" xml:"first_result,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelRatesLiteIncrUpdateResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRatesLiteIncrUpdateResultSet)
	},
}

// GetTaobaoXhotelRatesLiteIncrUpdateResultSet() 从对象池中获取TaobaoXhotelRatesLiteIncrUpdateResultSet
func GetTaobaoXhotelRatesLiteIncrUpdateResultSet() *TaobaoXhotelRatesLiteIncrUpdateResultSet {
	return poolTaobaoXhotelRatesLiteIncrUpdateResultSet.Get().(*TaobaoXhotelRatesLiteIncrUpdateResultSet)
}

// ReleaseTaobaoXhotelRatesLiteIncrUpdateResultSet 释放TaobaoXhotelRatesLiteIncrUpdateResultSet
func ReleaseTaobaoXhotelRatesLiteIncrUpdateResultSet(v *TaobaoXhotelRatesLiteIncrUpdateResultSet) {
	v.FirstResult = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTaobaoXhotelRatesLiteIncrUpdateResultSet.Put(v)
}
