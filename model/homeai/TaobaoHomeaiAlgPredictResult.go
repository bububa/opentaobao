package homeai

import (
	"sync"
)

// TaobaoHomeaiAlgPredictResult 结构体
type TaobaoHomeaiAlgPredictResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// errormsg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// data
	Data *FeatureWallSuggestionDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoHomeaiAlgPredictResult = sync.Pool{
	New: func() any {
		return new(TaobaoHomeaiAlgPredictResult)
	},
}

// GetTaobaoHomeaiAlgPredictResult() 从对象池中获取TaobaoHomeaiAlgPredictResult
func GetTaobaoHomeaiAlgPredictResult() *TaobaoHomeaiAlgPredictResult {
	return poolTaobaoHomeaiAlgPredictResult.Get().(*TaobaoHomeaiAlgPredictResult)
}

// ReleaseTaobaoHomeaiAlgPredictResult 释放TaobaoHomeaiAlgPredictResult
func ReleaseTaobaoHomeaiAlgPredictResult(v *TaobaoHomeaiAlgPredictResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = nil
	v.Success = false
	poolTaobaoHomeaiAlgPredictResult.Put(v)
}
