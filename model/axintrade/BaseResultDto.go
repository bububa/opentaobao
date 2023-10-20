package axintrade

import (
	"sync"
)

// BaseResultDto 结构体
type BaseResultDto struct {
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 信息提示
	InfoMsg string `json:"info_msg,omitempty" xml:"info_msg,omitempty"`
	// 是否处理审核结果成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否需要重试
	NeedRetry bool `json:"need_retry,omitempty" xml:"need_retry,omitempty"`
	// 是否调用成功
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
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.InfoMsg = ""
	v.Data = false
	v.NeedRetry = false
	v.Success = false
	poolBaseResultDto.Put(v)
}
