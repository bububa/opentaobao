package flight

import (
	"sync"
)

// AlitripAgentCoordinateProcessResult 结构体
type AlitripAgentCoordinateProcessResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentCoordinateProcessResult = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateProcessResult)
	},
}

// GetAlitripAgentCoordinateProcessResult() 从对象池中获取AlitripAgentCoordinateProcessResult
func GetAlitripAgentCoordinateProcessResult() *AlitripAgentCoordinateProcessResult {
	return poolAlitripAgentCoordinateProcessResult.Get().(*AlitripAgentCoordinateProcessResult)
}

// ReleaseAlitripAgentCoordinateProcessResult 释放AlitripAgentCoordinateProcessResult
func ReleaseAlitripAgentCoordinateProcessResult(v *AlitripAgentCoordinateProcessResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentCoordinateProcessResult.Put(v)
}
