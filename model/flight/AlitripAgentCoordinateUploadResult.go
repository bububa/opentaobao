package flight

import (
	"sync"
)

// AlitripAgentCoordinateUploadResult 结构体
type AlitripAgentCoordinateUploadResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 附件地址
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentCoordinateUploadResult = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateUploadResult)
	},
}

// GetAlitripAgentCoordinateUploadResult() 从对象池中获取AlitripAgentCoordinateUploadResult
func GetAlitripAgentCoordinateUploadResult() *AlitripAgentCoordinateUploadResult {
	return poolAlitripAgentCoordinateUploadResult.Get().(*AlitripAgentCoordinateUploadResult)
}

// ReleaseAlitripAgentCoordinateUploadResult 释放AlitripAgentCoordinateUploadResult
func ReleaseAlitripAgentCoordinateUploadResult(v *AlitripAgentCoordinateUploadResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = ""
	v.Success = false
	poolAlitripAgentCoordinateUploadResult.Put(v)
}
