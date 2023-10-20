package flight

import (
	"sync"
)

// AlitripAgentCoordinateProcessingResult 结构体
type AlitripAgentCoordinateProcessingResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentCoordinateProcessingResult = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateProcessingResult)
	},
}

// GetAlitripAgentCoordinateProcessingResult() 从对象池中获取AlitripAgentCoordinateProcessingResult
func GetAlitripAgentCoordinateProcessingResult() *AlitripAgentCoordinateProcessingResult {
	return poolAlitripAgentCoordinateProcessingResult.Get().(*AlitripAgentCoordinateProcessingResult)
}

// ReleaseAlitripAgentCoordinateProcessingResult 释放AlitripAgentCoordinateProcessingResult
func ReleaseAlitripAgentCoordinateProcessingResult(v *AlitripAgentCoordinateProcessingResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentCoordinateProcessingResult.Put(v)
}
