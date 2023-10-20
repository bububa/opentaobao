package alicom

import (
	"sync"
)

// AlibabaAliqinTaNumberSinglecallbyttsResult 结构体
type AlibabaAliqinTaNumberSinglecallbyttsResult struct {
	// 返回值
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 成功，失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinTaNumberSinglecallbyttsResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTaNumberSinglecallbyttsResult)
	},
}

// GetAlibabaAliqinTaNumberSinglecallbyttsResult() 从对象池中获取AlibabaAliqinTaNumberSinglecallbyttsResult
func GetAlibabaAliqinTaNumberSinglecallbyttsResult() *AlibabaAliqinTaNumberSinglecallbyttsResult {
	return poolAlibabaAliqinTaNumberSinglecallbyttsResult.Get().(*AlibabaAliqinTaNumberSinglecallbyttsResult)
}

// ReleaseAlibabaAliqinTaNumberSinglecallbyttsResult 释放AlibabaAliqinTaNumberSinglecallbyttsResult
func ReleaseAlibabaAliqinTaNumberSinglecallbyttsResult(v *AlibabaAliqinTaNumberSinglecallbyttsResult) {
	v.Model = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinTaNumberSinglecallbyttsResult.Put(v)
}
