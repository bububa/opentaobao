package idle

import (
	"sync"
)

// CarConsignmentResult 结构体
type CarConsignmentResult struct {
	// 错误信息
	PerformError string `json:"perform_error,omitempty" xml:"perform_error,omitempty"`
	// 请求成功
	TransformSuccess bool `json:"transform_success,omitempty" xml:"transform_success,omitempty"`
}

var poolCarConsignmentResult = sync.Pool{
	New: func() any {
		return new(CarConsignmentResult)
	},
}

// GetCarConsignmentResult() 从对象池中获取CarConsignmentResult
func GetCarConsignmentResult() *CarConsignmentResult {
	return poolCarConsignmentResult.Get().(*CarConsignmentResult)
}

// ReleaseCarConsignmentResult 释放CarConsignmentResult
func ReleaseCarConsignmentResult(v *CarConsignmentResult) {
	v.PerformError = ""
	v.TransformSuccess = false
	poolCarConsignmentResult.Put(v)
}
