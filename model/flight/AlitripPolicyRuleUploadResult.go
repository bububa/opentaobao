package flight

import (
	"sync"
)

// AlitripPolicyRuleUploadResult 结构体
type AlitripPolicyRuleUploadResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 创建&amp;删除结果参数
	Data *PolicyCreateResponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyRuleUploadResult = sync.Pool{
	New: func() any {
		return new(AlitripPolicyRuleUploadResult)
	},
}

// GetAlitripPolicyRuleUploadResult() 从对象池中获取AlitripPolicyRuleUploadResult
func GetAlitripPolicyRuleUploadResult() *AlitripPolicyRuleUploadResult {
	return poolAlitripPolicyRuleUploadResult.Get().(*AlitripPolicyRuleUploadResult)
}

// ReleaseAlitripPolicyRuleUploadResult 释放AlitripPolicyRuleUploadResult
func ReleaseAlitripPolicyRuleUploadResult(v *AlitripPolicyRuleUploadResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripPolicyRuleUploadResult.Put(v)
}
