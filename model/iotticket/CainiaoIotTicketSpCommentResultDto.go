package iotticket

import (
	"sync"
)

// CainiaoIotTicketSpCommentResultDto 结构体
type CainiaoIotTicketSpCommentResultDto struct {
	// 异常描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoIotTicketSpCommentResultDto = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpCommentResultDto)
	},
}

// GetCainiaoIotTicketSpCommentResultDto() 从对象池中获取CainiaoIotTicketSpCommentResultDto
func GetCainiaoIotTicketSpCommentResultDto() *CainiaoIotTicketSpCommentResultDto {
	return poolCainiaoIotTicketSpCommentResultDto.Get().(*CainiaoIotTicketSpCommentResultDto)
}

// ReleaseCainiaoIotTicketSpCommentResultDto 释放CainiaoIotTicketSpCommentResultDto
func ReleaseCainiaoIotTicketSpCommentResultDto(v *CainiaoIotTicketSpCommentResultDto) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Success = false
	poolCainiaoIotTicketSpCommentResultDto.Put(v)
}
