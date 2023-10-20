package lsttrade

import (
	"sync"
)

// AlibabaLstTradeOrderGetResultDto 结构体
type AlibabaLstTradeOrderGetResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回模型
	Content *LstTopOrderDto `json:"content,omitempty" xml:"content,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstTradeOrderGetResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeOrderGetResultDto)
	},
}

// GetAlibabaLstTradeOrderGetResultDto() 从对象池中获取AlibabaLstTradeOrderGetResultDto
func GetAlibabaLstTradeOrderGetResultDto() *AlibabaLstTradeOrderGetResultDto {
	return poolAlibabaLstTradeOrderGetResultDto.Get().(*AlibabaLstTradeOrderGetResultDto)
}

// ReleaseAlibabaLstTradeOrderGetResultDto 释放AlibabaLstTradeOrderGetResultDto
func ReleaseAlibabaLstTradeOrderGetResultDto(v *AlibabaLstTradeOrderGetResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolAlibabaLstTradeOrderGetResultDto.Put(v)
}
