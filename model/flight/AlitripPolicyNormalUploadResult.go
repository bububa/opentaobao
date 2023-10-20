package flight

import (
	"sync"
)

// AlitripPolicyNormalUploadResult 结构体
type AlitripPolicyNormalUploadResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 创建&amp;删除结果参数
	Data *PolicyCreateResponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyNormalUploadResult = sync.Pool{
	New: func() any {
		return new(AlitripPolicyNormalUploadResult)
	},
}

// GetAlitripPolicyNormalUploadResult() 从对象池中获取AlitripPolicyNormalUploadResult
func GetAlitripPolicyNormalUploadResult() *AlitripPolicyNormalUploadResult {
	return poolAlitripPolicyNormalUploadResult.Get().(*AlitripPolicyNormalUploadResult)
}

// ReleaseAlitripPolicyNormalUploadResult 释放AlitripPolicyNormalUploadResult
func ReleaseAlitripPolicyNormalUploadResult(v *AlitripPolicyNormalUploadResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripPolicyNormalUploadResult.Put(v)
}
