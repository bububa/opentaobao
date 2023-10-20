package sungari

import (
	"sync"
)

// TaobaoSungariDisposeQueryResult 结构体
type TaobaoSungariDisposeQueryResult struct {
	// data
	List []DisposeResultVo `json:"list,omitempty" xml:"list>dispose_result_vo,omitempty"`
	// 提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 服务是否调用成功，1：成功 2：失败 11：重复提交 其他：失败
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolTaobaoSungariDisposeQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoSungariDisposeQueryResult)
	},
}

// GetTaobaoSungariDisposeQueryResult() 从对象池中获取TaobaoSungariDisposeQueryResult
func GetTaobaoSungariDisposeQueryResult() *TaobaoSungariDisposeQueryResult {
	return poolTaobaoSungariDisposeQueryResult.Get().(*TaobaoSungariDisposeQueryResult)
}

// ReleaseTaobaoSungariDisposeQueryResult 释放TaobaoSungariDisposeQueryResult
func ReleaseTaobaoSungariDisposeQueryResult(v *TaobaoSungariDisposeQueryResult) {
	v.List = v.List[:0]
	v.Message = ""
	v.ResultCode = 0
	poolTaobaoSungariDisposeQueryResult.Put(v)
}
