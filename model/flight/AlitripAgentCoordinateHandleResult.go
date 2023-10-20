package flight

import (
	"sync"
)

// AlitripAgentCoordinateHandleResult 结构体
type AlitripAgentCoordinateHandleResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentCoordinateHandleResult = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateHandleResult)
	},
}

// GetAlitripAgentCoordinateHandleResult() 从对象池中获取AlitripAgentCoordinateHandleResult
func GetAlitripAgentCoordinateHandleResult() *AlitripAgentCoordinateHandleResult {
	return poolAlitripAgentCoordinateHandleResult.Get().(*AlitripAgentCoordinateHandleResult)
}

// ReleaseAlitripAgentCoordinateHandleResult 释放AlitripAgentCoordinateHandleResult
func ReleaseAlitripAgentCoordinateHandleResult(v *AlitripAgentCoordinateHandleResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentCoordinateHandleResult.Put(v)
}
