package iotticket

import (
	"sync"
)

// CainiaoIotTicketSpMailSignUploadResultDto 结构体
type CainiaoIotTicketSpMailSignUploadResultDto struct {
	// 异常描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 异常编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoIotTicketSpMailSignUploadResultDto = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpMailSignUploadResultDto)
	},
}

// GetCainiaoIotTicketSpMailSignUploadResultDto() 从对象池中获取CainiaoIotTicketSpMailSignUploadResultDto
func GetCainiaoIotTicketSpMailSignUploadResultDto() *CainiaoIotTicketSpMailSignUploadResultDto {
	return poolCainiaoIotTicketSpMailSignUploadResultDto.Get().(*CainiaoIotTicketSpMailSignUploadResultDto)
}

// ReleaseCainiaoIotTicketSpMailSignUploadResultDto 释放CainiaoIotTicketSpMailSignUploadResultDto
func ReleaseCainiaoIotTicketSpMailSignUploadResultDto(v *CainiaoIotTicketSpMailSignUploadResultDto) {
	v.ErrorDesc = ""
	v.ErrorCode = ""
	v.Success = false
	poolCainiaoIotTicketSpMailSignUploadResultDto.Put(v)
}
