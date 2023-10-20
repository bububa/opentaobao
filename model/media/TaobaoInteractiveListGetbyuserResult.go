package media

import (
	"sync"
)

// TaobaoInteractiveListGetbyuserResult 结构体
type TaobaoInteractiveListGetbyuserResult struct {
	// 错误
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// model
	Model *PageQueryResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoInteractiveListGetbyuserResult = sync.Pool{
	New: func() any {
		return new(TaobaoInteractiveListGetbyuserResult)
	},
}

// GetTaobaoInteractiveListGetbyuserResult() 从对象池中获取TaobaoInteractiveListGetbyuserResult
func GetTaobaoInteractiveListGetbyuserResult() *TaobaoInteractiveListGetbyuserResult {
	return poolTaobaoInteractiveListGetbyuserResult.Get().(*TaobaoInteractiveListGetbyuserResult)
}

// ReleaseTaobaoInteractiveListGetbyuserResult 释放TaobaoInteractiveListGetbyuserResult
func ReleaseTaobaoInteractiveListGetbyuserResult(v *TaobaoInteractiveListGetbyuserResult) {
	v.ResultCode = nil
	v.Model = nil
	v.Success = false
	poolTaobaoInteractiveListGetbyuserResult.Put(v)
}
