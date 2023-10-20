package jym

import (
	"sync"
)

// ResultDto 结构体
type ResultDto struct {
	// 调用错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 调用错误码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 是否调用成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 业务对象
	Result *GoodsResultDto `json:"result,omitempty" xml:"result,omitempty"`
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
	v.ExtraErrMsg = ""
	v.StateCode = ""
	v.Success = ""
	v.Result = nil
	poolResultDto.Put(v)
}
