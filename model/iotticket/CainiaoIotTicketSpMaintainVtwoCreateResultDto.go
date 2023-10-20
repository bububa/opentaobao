package iotticket

import (
	"sync"
)

// CainiaoIotTicketSpMaintainVtwoCreateResultDto 结构体
type CainiaoIotTicketSpMaintainVtwoCreateResultDto struct {
	// 异常描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoIotTicketSpMaintainVtwoCreateResultDto = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpMaintainVtwoCreateResultDto)
	},
}

// GetCainiaoIotTicketSpMaintainVtwoCreateResultDto() 从对象池中获取CainiaoIotTicketSpMaintainVtwoCreateResultDto
func GetCainiaoIotTicketSpMaintainVtwoCreateResultDto() *CainiaoIotTicketSpMaintainVtwoCreateResultDto {
	return poolCainiaoIotTicketSpMaintainVtwoCreateResultDto.Get().(*CainiaoIotTicketSpMaintainVtwoCreateResultDto)
}

// ReleaseCainiaoIotTicketSpMaintainVtwoCreateResultDto 释放CainiaoIotTicketSpMaintainVtwoCreateResultDto
func ReleaseCainiaoIotTicketSpMaintainVtwoCreateResultDto(v *CainiaoIotTicketSpMaintainVtwoCreateResultDto) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Success = false
	poolCainiaoIotTicketSpMaintainVtwoCreateResultDto.Put(v)
}
