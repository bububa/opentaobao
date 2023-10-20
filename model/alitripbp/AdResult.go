package alitripbp

import (
	"sync"
)

// AdResult 结构体
type AdResult struct {
	// 1
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 查询结果
	Model *ChannelExamineResultDto `json:"model,omitempty" xml:"model,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAdResult = sync.Pool{
	New: func() any {
		return new(AdResult)
	},
}

// GetAdResult() 从对象池中获取AdResult
func GetAdResult() *AdResult {
	return poolAdResult.Get().(*AdResult)
}

// ReleaseAdResult 释放AdResult
func ReleaseAdResult(v *AdResult) {
	v.Msg = ""
	v.Code = ""
	v.Model = nil
	v.Success = false
	poolAdResult.Put(v)
}
