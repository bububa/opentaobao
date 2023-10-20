package flight

import (
	"sync"
)

// AlitripPolicyNormalCompressionUploadResultDto 结构体
type AlitripPolicyNormalCompressionUploadResultDto struct {
	// 任务失败错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 任务失败错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyNormalCompressionUploadResultDto = sync.Pool{
	New: func() any {
		return new(AlitripPolicyNormalCompressionUploadResultDto)
	},
}

// GetAlitripPolicyNormalCompressionUploadResultDto() 从对象池中获取AlitripPolicyNormalCompressionUploadResultDto
func GetAlitripPolicyNormalCompressionUploadResultDto() *AlitripPolicyNormalCompressionUploadResultDto {
	return poolAlitripPolicyNormalCompressionUploadResultDto.Get().(*AlitripPolicyNormalCompressionUploadResultDto)
}

// ReleaseAlitripPolicyNormalCompressionUploadResultDto 释放AlitripPolicyNormalCompressionUploadResultDto
func ReleaseAlitripPolicyNormalCompressionUploadResultDto(v *AlitripPolicyNormalCompressionUploadResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = 0
	v.Success = false
	poolAlitripPolicyNormalCompressionUploadResultDto.Put(v)
}
