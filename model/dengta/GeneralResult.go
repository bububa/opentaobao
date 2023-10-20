package dengta

import (
	"sync"
)

// GeneralResult 结构体
type GeneralResult struct {
	// 请求标识
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Value bool `json:"value,omitempty" xml:"value,omitempty"`
}

var poolGeneralResult = sync.Pool{
	New: func() any {
		return new(GeneralResult)
	},
}

// GetGeneralResult() 从对象池中获取GeneralResult
func GetGeneralResult() *GeneralResult {
	return poolGeneralResult.Get().(*GeneralResult)
}

// ReleaseGeneralResult 释放GeneralResult
func ReleaseGeneralResult(v *GeneralResult) {
	v.TraceId = ""
	v.Code = ""
	v.Value = false
	poolGeneralResult.Put(v)
}
