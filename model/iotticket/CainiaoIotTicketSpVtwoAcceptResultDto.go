package iotticket

import (
	"sync"
)

// CainiaoIotTicketSpVtwoAcceptResultDto 结构体
type CainiaoIotTicketSpVtwoAcceptResultDto struct {
	// 异常描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoIotTicketSpVtwoAcceptResultDto = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpVtwoAcceptResultDto)
	},
}

// GetCainiaoIotTicketSpVtwoAcceptResultDto() 从对象池中获取CainiaoIotTicketSpVtwoAcceptResultDto
func GetCainiaoIotTicketSpVtwoAcceptResultDto() *CainiaoIotTicketSpVtwoAcceptResultDto {
	return poolCainiaoIotTicketSpVtwoAcceptResultDto.Get().(*CainiaoIotTicketSpVtwoAcceptResultDto)
}

// ReleaseCainiaoIotTicketSpVtwoAcceptResultDto 释放CainiaoIotTicketSpVtwoAcceptResultDto
func ReleaseCainiaoIotTicketSpVtwoAcceptResultDto(v *CainiaoIotTicketSpVtwoAcceptResultDto) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Success = false
	poolCainiaoIotTicketSpVtwoAcceptResultDto.Put(v)
}
