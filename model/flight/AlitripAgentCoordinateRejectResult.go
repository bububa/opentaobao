package flight

import (
	"sync"
)

// AlitripAgentCoordinateRejectResult 结构体
type AlitripAgentCoordinateRejectResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentCoordinateRejectResult = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateRejectResult)
	},
}

// GetAlitripAgentCoordinateRejectResult() 从对象池中获取AlitripAgentCoordinateRejectResult
func GetAlitripAgentCoordinateRejectResult() *AlitripAgentCoordinateRejectResult {
	return poolAlitripAgentCoordinateRejectResult.Get().(*AlitripAgentCoordinateRejectResult)
}

// ReleaseAlitripAgentCoordinateRejectResult 释放AlitripAgentCoordinateRejectResult
func ReleaseAlitripAgentCoordinateRejectResult(v *AlitripAgentCoordinateRejectResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentCoordinateRejectResult.Put(v)
}
