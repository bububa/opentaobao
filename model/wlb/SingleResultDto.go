package wlb

import (
	"sync"
)

// SingleResultDto 结构体
type SingleResultDto struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Flag bool `json:"flag,omitempty" xml:"flag,omitempty"`
}

var poolSingleResultDto = sync.Pool{
	New: func() any {
		return new(SingleResultDto)
	},
}

// GetSingleResultDto() 从对象池中获取SingleResultDto
func GetSingleResultDto() *SingleResultDto {
	return poolSingleResultDto.Get().(*SingleResultDto)
}

// ReleaseSingleResultDto 释放SingleResultDto
func ReleaseSingleResultDto(v *SingleResultDto) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Flag = false
	poolSingleResultDto.Put(v)
}
