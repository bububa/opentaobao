package lsttrade

import (
	"sync"
)

// AlibabaLstTradeSellerOrderDetailQueryResultDto 结构体
type AlibabaLstTradeSellerOrderDetailQueryResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回值
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstTradeSellerOrderDetailQueryResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerOrderDetailQueryResultDto)
	},
}

// GetAlibabaLstTradeSellerOrderDetailQueryResultDto() 从对象池中获取AlibabaLstTradeSellerOrderDetailQueryResultDto
func GetAlibabaLstTradeSellerOrderDetailQueryResultDto() *AlibabaLstTradeSellerOrderDetailQueryResultDto {
	return poolAlibabaLstTradeSellerOrderDetailQueryResultDto.Get().(*AlibabaLstTradeSellerOrderDetailQueryResultDto)
}

// ReleaseAlibabaLstTradeSellerOrderDetailQueryResultDto 释放AlibabaLstTradeSellerOrderDetailQueryResultDto
func ReleaseAlibabaLstTradeSellerOrderDetailQueryResultDto(v *AlibabaLstTradeSellerOrderDetailQueryResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolAlibabaLstTradeSellerOrderDetailQueryResultDto.Put(v)
}
