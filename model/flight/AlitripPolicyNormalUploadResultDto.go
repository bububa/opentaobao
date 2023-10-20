package flight

import (
	"sync"
)

// AlitripPolicyNormalUploadResultDto 结构体
type AlitripPolicyNormalUploadResultDto struct {
	// 任务失败错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 任务失败错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyNormalUploadResultDto = sync.Pool{
	New: func() any {
		return new(AlitripPolicyNormalUploadResultDto)
	},
}

// GetAlitripPolicyNormalUploadResultDto() 从对象池中获取AlitripPolicyNormalUploadResultDto
func GetAlitripPolicyNormalUploadResultDto() *AlitripPolicyNormalUploadResultDto {
	return poolAlitripPolicyNormalUploadResultDto.Get().(*AlitripPolicyNormalUploadResultDto)
}

// ReleaseAlitripPolicyNormalUploadResultDto 释放AlitripPolicyNormalUploadResultDto
func ReleaseAlitripPolicyNormalUploadResultDto(v *AlitripPolicyNormalUploadResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = 0
	v.Success = false
	poolAlitripPolicyNormalUploadResultDto.Put(v)
}
