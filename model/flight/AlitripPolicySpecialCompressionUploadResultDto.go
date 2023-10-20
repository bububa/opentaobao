package flight

import (
	"sync"
)

// AlitripPolicySpecialCompressionUploadResultDto 结构体
type AlitripPolicySpecialCompressionUploadResultDto struct {
	// 任务失败错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 任务失败错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicySpecialCompressionUploadResultDto = sync.Pool{
	New: func() any {
		return new(AlitripPolicySpecialCompressionUploadResultDto)
	},
}

// GetAlitripPolicySpecialCompressionUploadResultDto() 从对象池中获取AlitripPolicySpecialCompressionUploadResultDto
func GetAlitripPolicySpecialCompressionUploadResultDto() *AlitripPolicySpecialCompressionUploadResultDto {
	return poolAlitripPolicySpecialCompressionUploadResultDto.Get().(*AlitripPolicySpecialCompressionUploadResultDto)
}

// ReleaseAlitripPolicySpecialCompressionUploadResultDto 释放AlitripPolicySpecialCompressionUploadResultDto
func ReleaseAlitripPolicySpecialCompressionUploadResultDto(v *AlitripPolicySpecialCompressionUploadResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = 0
	v.Success = false
	poolAlitripPolicySpecialCompressionUploadResultDto.Put(v)
}
