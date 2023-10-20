package logistic

import (
	"sync"
)

// SingleResultDto 结构体
type SingleResultDto struct {
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 调用码
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 业务返回结果
	Result *ExpressModifyAppointTopResponseDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否需要重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
	// 是否幂等
	IsIdempotent bool `json:"is_idempotent,omitempty" xml:"is_idempotent,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.ErrorDesc = ""
	v.TraceId = ""
	v.ErrorCode = ""
	v.Result = nil
	v.IsRetry = false
	v.IsIdempotent = false
	v.Success = false
	poolSingleResultDto.Put(v)
}
