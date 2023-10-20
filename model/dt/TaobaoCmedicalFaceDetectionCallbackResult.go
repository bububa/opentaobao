package dt

import (
	"sync"
)

// TaobaoCmedicalFaceDetectionCallbackResult 结构体
type TaobaoCmedicalFaceDetectionCallbackResult struct {
	// 错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功，true：成功，false：失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoCmedicalFaceDetectionCallbackResult = sync.Pool{
	New: func() any {
		return new(TaobaoCmedicalFaceDetectionCallbackResult)
	},
}

// GetTaobaoCmedicalFaceDetectionCallbackResult() 从对象池中获取TaobaoCmedicalFaceDetectionCallbackResult
func GetTaobaoCmedicalFaceDetectionCallbackResult() *TaobaoCmedicalFaceDetectionCallbackResult {
	return poolTaobaoCmedicalFaceDetectionCallbackResult.Get().(*TaobaoCmedicalFaceDetectionCallbackResult)
}

// ReleaseTaobaoCmedicalFaceDetectionCallbackResult 释放TaobaoCmedicalFaceDetectionCallbackResult
func ReleaseTaobaoCmedicalFaceDetectionCallbackResult(v *TaobaoCmedicalFaceDetectionCallbackResult) {
	v.Error = nil
	v.Success = false
	poolTaobaoCmedicalFaceDetectionCallbackResult.Put(v)
}
