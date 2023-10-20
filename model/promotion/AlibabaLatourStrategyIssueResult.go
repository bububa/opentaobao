package promotion

import (
	"sync"
)

// AlibabaLatourStrategyIssueResult 结构体
type AlibabaLatourStrategyIssueResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回结果
	Data *StrategyIssueResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLatourStrategyIssueResult = sync.Pool{
	New: func() any {
		return new(AlibabaLatourStrategyIssueResult)
	},
}

// GetAlibabaLatourStrategyIssueResult() 从对象池中获取AlibabaLatourStrategyIssueResult
func GetAlibabaLatourStrategyIssueResult() *AlibabaLatourStrategyIssueResult {
	return poolAlibabaLatourStrategyIssueResult.Get().(*AlibabaLatourStrategyIssueResult)
}

// ReleaseAlibabaLatourStrategyIssueResult 释放AlibabaLatourStrategyIssueResult
func ReleaseAlibabaLatourStrategyIssueResult(v *AlibabaLatourStrategyIssueResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = nil
	v.Success = false
	poolAlibabaLatourStrategyIssueResult.Put(v)
}
