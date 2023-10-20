package logistic

import (
	"sync"
)

// BaseResultDto 结构体
type BaseResultDto struct {
	// 请求错误信息
	OneErrorInfo *ErrorInfo `json:"one_error_info,omitempty" xml:"one_error_info,omitempty"`
	// 返回信息
	Module *ReachableServiceWaybillForTopResponseDto `json:"module,omitempty" xml:"module,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBaseResultDto = sync.Pool{
	New: func() any {
		return new(BaseResultDto)
	},
}

// GetBaseResultDto() 从对象池中获取BaseResultDto
func GetBaseResultDto() *BaseResultDto {
	return poolBaseResultDto.Get().(*BaseResultDto)
}

// ReleaseBaseResultDto 释放BaseResultDto
func ReleaseBaseResultDto(v *BaseResultDto) {
	v.OneErrorInfo = nil
	v.Module = nil
	v.Success = false
	poolBaseResultDto.Put(v)
}
