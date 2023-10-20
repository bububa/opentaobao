package flight

import (
	"sync"
)

// AlitripPolicyProcessResultDto 结构体
type AlitripPolicyProcessResultDto struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 创建&amp;删除结果参数
	Data *PolicyCreateResponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyProcessResultDto = sync.Pool{
	New: func() any {
		return new(AlitripPolicyProcessResultDto)
	},
}

// GetAlitripPolicyProcessResultDto() 从对象池中获取AlitripPolicyProcessResultDto
func GetAlitripPolicyProcessResultDto() *AlitripPolicyProcessResultDto {
	return poolAlitripPolicyProcessResultDto.Get().(*AlitripPolicyProcessResultDto)
}

// ReleaseAlitripPolicyProcessResultDto 释放AlitripPolicyProcessResultDto
func ReleaseAlitripPolicyProcessResultDto(v *AlitripPolicyProcessResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripPolicyProcessResultDto.Put(v)
}
