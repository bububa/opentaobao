package lstvending

import (
	"sync"
)

// AlibabaLstVendingCargospaceSaveResultDto 结构体
type AlibabaLstVendingCargospaceSaveResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 记录唯一标识
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolAlibabaLstVendingCargospaceSaveResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingCargospaceSaveResultDto)
	},
}

// GetAlibabaLstVendingCargospaceSaveResultDto() 从对象池中获取AlibabaLstVendingCargospaceSaveResultDto
func GetAlibabaLstVendingCargospaceSaveResultDto() *AlibabaLstVendingCargospaceSaveResultDto {
	return poolAlibabaLstVendingCargospaceSaveResultDto.Get().(*AlibabaLstVendingCargospaceSaveResultDto)
}

// ReleaseAlibabaLstVendingCargospaceSaveResultDto 释放AlibabaLstVendingCargospaceSaveResultDto
func ReleaseAlibabaLstVendingCargospaceSaveResultDto(v *AlibabaLstVendingCargospaceSaveResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Key = ""
	poolAlibabaLstVendingCargospaceSaveResultDto.Put(v)
}
