package dengta

import (
	"sync"
)

// ApiGeneralResult 结构体
type ApiGeneralResult struct {
	// 链路id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息
	Message *ReturnMessage `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Value bool `json:"value,omitempty" xml:"value,omitempty"`
}

var poolApiGeneralResult = sync.Pool{
	New: func() any {
		return new(ApiGeneralResult)
	},
}

// GetApiGeneralResult() 从对象池中获取ApiGeneralResult
func GetApiGeneralResult() *ApiGeneralResult {
	return poolApiGeneralResult.Get().(*ApiGeneralResult)
}

// ReleaseApiGeneralResult 释放ApiGeneralResult
func ReleaseApiGeneralResult(v *ApiGeneralResult) {
	v.TraceId = ""
	v.Message = nil
	v.Code = 0
	v.Value = false
	poolApiGeneralResult.Put(v)
}
