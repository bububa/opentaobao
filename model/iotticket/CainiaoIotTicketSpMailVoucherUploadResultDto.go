package iotticket

import (
	"sync"
)

// CainiaoIotTicketSpMailVoucherUploadResultDto 结构体
type CainiaoIotTicketSpMailVoucherUploadResultDto struct {
	// 异常描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 异常编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoIotTicketSpMailVoucherUploadResultDto = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpMailVoucherUploadResultDto)
	},
}

// GetCainiaoIotTicketSpMailVoucherUploadResultDto() 从对象池中获取CainiaoIotTicketSpMailVoucherUploadResultDto
func GetCainiaoIotTicketSpMailVoucherUploadResultDto() *CainiaoIotTicketSpMailVoucherUploadResultDto {
	return poolCainiaoIotTicketSpMailVoucherUploadResultDto.Get().(*CainiaoIotTicketSpMailVoucherUploadResultDto)
}

// ReleaseCainiaoIotTicketSpMailVoucherUploadResultDto 释放CainiaoIotTicketSpMailVoucherUploadResultDto
func ReleaseCainiaoIotTicketSpMailVoucherUploadResultDto(v *CainiaoIotTicketSpMailVoucherUploadResultDto) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Success = false
	poolCainiaoIotTicketSpMailVoucherUploadResultDto.Put(v)
}
