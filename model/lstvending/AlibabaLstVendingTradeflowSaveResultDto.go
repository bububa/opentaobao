package lstvending

import (
	"sync"
)

// AlibabaLstVendingTradeflowSaveResultDto 结构体
type AlibabaLstVendingTradeflowSaveResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行失败记录ID
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolAlibabaLstVendingTradeflowSaveResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingTradeflowSaveResultDto)
	},
}

// GetAlibabaLstVendingTradeflowSaveResultDto() 从对象池中获取AlibabaLstVendingTradeflowSaveResultDto
func GetAlibabaLstVendingTradeflowSaveResultDto() *AlibabaLstVendingTradeflowSaveResultDto {
	return poolAlibabaLstVendingTradeflowSaveResultDto.Get().(*AlibabaLstVendingTradeflowSaveResultDto)
}

// ReleaseAlibabaLstVendingTradeflowSaveResultDto 释放AlibabaLstVendingTradeflowSaveResultDto
func ReleaseAlibabaLstVendingTradeflowSaveResultDto(v *AlibabaLstVendingTradeflowSaveResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Key = ""
	poolAlibabaLstVendingTradeflowSaveResultDto.Put(v)
}
