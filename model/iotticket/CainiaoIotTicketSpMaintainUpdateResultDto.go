package iotticket

import (
	"sync"
)

// CainiaoIotTicketSpMaintainUpdateResultDto 结构体
type CainiaoIotTicketSpMaintainUpdateResultDto struct {
	// 异常描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 异常编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoIotTicketSpMaintainUpdateResultDto = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpMaintainUpdateResultDto)
	},
}

// GetCainiaoIotTicketSpMaintainUpdateResultDto() 从对象池中获取CainiaoIotTicketSpMaintainUpdateResultDto
func GetCainiaoIotTicketSpMaintainUpdateResultDto() *CainiaoIotTicketSpMaintainUpdateResultDto {
	return poolCainiaoIotTicketSpMaintainUpdateResultDto.Get().(*CainiaoIotTicketSpMaintainUpdateResultDto)
}

// ReleaseCainiaoIotTicketSpMaintainUpdateResultDto 释放CainiaoIotTicketSpMaintainUpdateResultDto
func ReleaseCainiaoIotTicketSpMaintainUpdateResultDto(v *CainiaoIotTicketSpMaintainUpdateResultDto) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Success = false
	poolCainiaoIotTicketSpMaintainUpdateResultDto.Put(v)
}
