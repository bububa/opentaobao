package lstvending

import (
	"sync"
)

// AlibabaLstVendingTradeflowQueryResultDto 结构体
type AlibabaLstVendingTradeflowQueryResultDto struct {
	// 交易流水记录
	ModuleList []VendingTradeFlowDto `json:"module_list,omitempty" xml:"module_list>vending_trade_flow_dto,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstVendingTradeflowQueryResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingTradeflowQueryResultDto)
	},
}

// GetAlibabaLstVendingTradeflowQueryResultDto() 从对象池中获取AlibabaLstVendingTradeflowQueryResultDto
func GetAlibabaLstVendingTradeflowQueryResultDto() *AlibabaLstVendingTradeflowQueryResultDto {
	return poolAlibabaLstVendingTradeflowQueryResultDto.Get().(*AlibabaLstVendingTradeflowQueryResultDto)
}

// ReleaseAlibabaLstVendingTradeflowQueryResultDto 释放AlibabaLstVendingTradeflowQueryResultDto
func ReleaseAlibabaLstVendingTradeflowQueryResultDto(v *AlibabaLstVendingTradeflowQueryResultDto) {
	v.ModuleList = v.ModuleList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaLstVendingTradeflowQueryResultDto.Put(v)
}
