package omniorder

import (
	"sync"
)

// TaobaoOmniorderStorecollectQueryResult 结构体
type TaobaoOmniorderStorecollectQueryResult struct {
	// 0表示码可用，其余值表示码不可用
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 码状态附加信息，若码可用则此处为空
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *StoreCollectQueryOrderResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoOmniorderStorecollectQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderStorecollectQueryResult)
	},
}

// GetTaobaoOmniorderStorecollectQueryResult() 从对象池中获取TaobaoOmniorderStorecollectQueryResult
func GetTaobaoOmniorderStorecollectQueryResult() *TaobaoOmniorderStorecollectQueryResult {
	return poolTaobaoOmniorderStorecollectQueryResult.Get().(*TaobaoOmniorderStorecollectQueryResult)
}

// ReleaseTaobaoOmniorderStorecollectQueryResult 释放TaobaoOmniorderStorecollectQueryResult
func ReleaseTaobaoOmniorderStorecollectQueryResult(v *TaobaoOmniorderStorecollectQueryResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	poolTaobaoOmniorderStorecollectQueryResult.Put(v)
}
