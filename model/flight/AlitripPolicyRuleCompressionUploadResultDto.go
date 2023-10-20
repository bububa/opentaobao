package flight

import (
	"sync"
)

// AlitripPolicyRuleCompressionUploadResultDto 结构体
type AlitripPolicyRuleCompressionUploadResultDto struct {
	// 任务失败错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 任务失败错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyRuleCompressionUploadResultDto = sync.Pool{
	New: func() any {
		return new(AlitripPolicyRuleCompressionUploadResultDto)
	},
}

// GetAlitripPolicyRuleCompressionUploadResultDto() 从对象池中获取AlitripPolicyRuleCompressionUploadResultDto
func GetAlitripPolicyRuleCompressionUploadResultDto() *AlitripPolicyRuleCompressionUploadResultDto {
	return poolAlitripPolicyRuleCompressionUploadResultDto.Get().(*AlitripPolicyRuleCompressionUploadResultDto)
}

// ReleaseAlitripPolicyRuleCompressionUploadResultDto 释放AlitripPolicyRuleCompressionUploadResultDto
func ReleaseAlitripPolicyRuleCompressionUploadResultDto(v *AlitripPolicyRuleCompressionUploadResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = 0
	v.Success = false
	poolAlitripPolicyRuleCompressionUploadResultDto.Put(v)
}
