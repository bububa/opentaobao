package feedflow

import (
	"sync"
)

// ResultDto 结构体
type ResultDto struct {
	// 返回说明信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 账户信息
	Result *AccountDto `json:"result,omitempty" xml:"result,omitempty"`
	// 调用是否成功
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
	v.Message = ""
	v.Result = nil
	v.Success = false
	poolResultDto.Put(v)
}
