package flight

import (
	"sync"
)

// AlitripPolicySpecialUploadResult 结构体
type AlitripPolicySpecialUploadResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 创建&amp;删除结果参数
	Data *PolicyCreateResponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicySpecialUploadResult = sync.Pool{
	New: func() any {
		return new(AlitripPolicySpecialUploadResult)
	},
}

// GetAlitripPolicySpecialUploadResult() 从对象池中获取AlitripPolicySpecialUploadResult
func GetAlitripPolicySpecialUploadResult() *AlitripPolicySpecialUploadResult {
	return poolAlitripPolicySpecialUploadResult.Get().(*AlitripPolicySpecialUploadResult)
}

// ReleaseAlitripPolicySpecialUploadResult 释放AlitripPolicySpecialUploadResult
func ReleaseAlitripPolicySpecialUploadResult(v *AlitripPolicySpecialUploadResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripPolicySpecialUploadResult.Put(v)
}
