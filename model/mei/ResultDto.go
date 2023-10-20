package mei

import (
	"sync"
)

// ResultDto 结构体
type ResultDto struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// total
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// result
	Result *MemberAccountDto `json:"result,omitempty" xml:"result,omitempty"`
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
	v.Code = ""
	v.Msg = ""
	v.Total = 0
	v.Result = nil
	poolResultDto.Put(v)
}
