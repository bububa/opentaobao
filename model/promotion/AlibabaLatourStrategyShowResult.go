package promotion

import (
	"sync"
)

// AlibabaLatourStrategyShowResult 结构体
type AlibabaLatourStrategyShowResult struct {
	// 错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回结果
	Data *StrategyShowResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLatourStrategyShowResult = sync.Pool{
	New: func() any {
		return new(AlibabaLatourStrategyShowResult)
	},
}

// GetAlibabaLatourStrategyShowResult() 从对象池中获取AlibabaLatourStrategyShowResult
func GetAlibabaLatourStrategyShowResult() *AlibabaLatourStrategyShowResult {
	return poolAlibabaLatourStrategyShowResult.Get().(*AlibabaLatourStrategyShowResult)
}

// ReleaseAlibabaLatourStrategyShowResult 释放AlibabaLatourStrategyShowResult
func ReleaseAlibabaLatourStrategyShowResult(v *AlibabaLatourStrategyShowResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaLatourStrategyShowResult.Put(v)
}
