package lstvending

import (
	"sync"
)

// AlibabaLstVendingGoodsSaveResultDto 结构体
type AlibabaLstVendingGoodsSaveResultDto struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 记录唯一值
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolAlibabaLstVendingGoodsSaveResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingGoodsSaveResultDto)
	},
}

// GetAlibabaLstVendingGoodsSaveResultDto() 从对象池中获取AlibabaLstVendingGoodsSaveResultDto
func GetAlibabaLstVendingGoodsSaveResultDto() *AlibabaLstVendingGoodsSaveResultDto {
	return poolAlibabaLstVendingGoodsSaveResultDto.Get().(*AlibabaLstVendingGoodsSaveResultDto)
}

// ReleaseAlibabaLstVendingGoodsSaveResultDto 释放AlibabaLstVendingGoodsSaveResultDto
func ReleaseAlibabaLstVendingGoodsSaveResultDto(v *AlibabaLstVendingGoodsSaveResultDto) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Key = ""
	poolAlibabaLstVendingGoodsSaveResultDto.Put(v)
}
