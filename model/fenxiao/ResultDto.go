package fenxiao

import (
	"sync"
)

// ResultDto 结构体
type ResultDto struct {
	// 库存数量
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultDto = sync.Pool{
	New: func() any {
		return new(ResultDto)
	},
}

// GetResultDto() 从对象池中获取ResultDto
func GetResultDto() *ResultDto {
	return poolResultDto.Get().(*ResultDto)
}

// ReleaseResultDto 释放ResultDto
func ReleaseResultDto(v *ResultDto) {
	v.Module = ""
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolResultDto.Put(v)
}
