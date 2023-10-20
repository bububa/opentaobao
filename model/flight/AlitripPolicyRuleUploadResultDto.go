package flight

import (
	"sync"
)

// AlitripPolicyRuleUploadResultDto 结构体
type AlitripPolicyRuleUploadResultDto struct {
	// 任务失败错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 任务失败错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyRuleUploadResultDto = sync.Pool{
	New: func() any {
		return new(AlitripPolicyRuleUploadResultDto)
	},
}

// GetAlitripPolicyRuleUploadResultDto() 从对象池中获取AlitripPolicyRuleUploadResultDto
func GetAlitripPolicyRuleUploadResultDto() *AlitripPolicyRuleUploadResultDto {
	return poolAlitripPolicyRuleUploadResultDto.Get().(*AlitripPolicyRuleUploadResultDto)
}

// ReleaseAlitripPolicyRuleUploadResultDto 释放AlitripPolicyRuleUploadResultDto
func ReleaseAlitripPolicyRuleUploadResultDto(v *AlitripPolicyRuleUploadResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = 0
	v.Success = false
	poolAlitripPolicyRuleUploadResultDto.Put(v)
}
