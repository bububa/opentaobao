package flight

import (
	"sync"
)

// AlitripPolicySpecialUploadResultDto 结构体
type AlitripPolicySpecialUploadResultDto struct {
	// 任务失败错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 任务失败错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicySpecialUploadResultDto = sync.Pool{
	New: func() any {
		return new(AlitripPolicySpecialUploadResultDto)
	},
}

// GetAlitripPolicySpecialUploadResultDto() 从对象池中获取AlitripPolicySpecialUploadResultDto
func GetAlitripPolicySpecialUploadResultDto() *AlitripPolicySpecialUploadResultDto {
	return poolAlitripPolicySpecialUploadResultDto.Get().(*AlitripPolicySpecialUploadResultDto)
}

// ReleaseAlitripPolicySpecialUploadResultDto 释放AlitripPolicySpecialUploadResultDto
func ReleaseAlitripPolicySpecialUploadResultDto(v *AlitripPolicySpecialUploadResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = 0
	v.Success = false
	poolAlitripPolicySpecialUploadResultDto.Put(v)
}
