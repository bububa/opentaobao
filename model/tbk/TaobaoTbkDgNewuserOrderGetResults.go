package tbk

import (
	"sync"
)

// TaobaoTbkDgNewuserOrderGetResults 结构体
type TaobaoTbkDgNewuserOrderGetResults struct {
	// data
	Data *TaobaoTbkDgNewuserOrderGetData `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoTbkDgNewuserOrderGetResults = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgNewuserOrderGetResults)
	},
}

// GetTaobaoTbkDgNewuserOrderGetResults() 从对象池中获取TaobaoTbkDgNewuserOrderGetResults
func GetTaobaoTbkDgNewuserOrderGetResults() *TaobaoTbkDgNewuserOrderGetResults {
	return poolTaobaoTbkDgNewuserOrderGetResults.Get().(*TaobaoTbkDgNewuserOrderGetResults)
}

// ReleaseTaobaoTbkDgNewuserOrderGetResults 释放TaobaoTbkDgNewuserOrderGetResults
func ReleaseTaobaoTbkDgNewuserOrderGetResults(v *TaobaoTbkDgNewuserOrderGetResults) {
	v.Data = nil
	poolTaobaoTbkDgNewuserOrderGetResults.Put(v)
}
