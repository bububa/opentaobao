package alicom

import (
	"sync"
)

// TopResultDto 结构体
type TopResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data *ProductActivityInfoResponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopResultDto = sync.Pool{
	New: func() any {
		return new(TopResultDto)
	},
}

// GetTopResultDto() 从对象池中获取TopResultDto
func GetTopResultDto() *TopResultDto {
	return poolTopResultDto.Get().(*TopResultDto)
}

// ReleaseTopResultDto 释放TopResultDto
func ReleaseTopResultDto(v *TopResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolTopResultDto.Put(v)
}
