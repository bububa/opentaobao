package flight

import (
	"sync"
)

// ResultDto 结构体
type ResultDto struct {
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 查询协同单详情返回内容
	Data *CaseResultDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// true
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
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolResultDto.Put(v)
}
