package lstmarketing

import (
	"sync"
)

// AlibabaLstMarketingQuerybyorderidResultDto 结构体
type AlibabaLstMarketingQuerybyorderidResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 订单实体
	Content *LstTopOrderDto `json:"content,omitempty" xml:"content,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstMarketingQuerybyorderidResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstMarketingQuerybyorderidResultDto)
	},
}

// GetAlibabaLstMarketingQuerybyorderidResultDto() 从对象池中获取AlibabaLstMarketingQuerybyorderidResultDto
func GetAlibabaLstMarketingQuerybyorderidResultDto() *AlibabaLstMarketingQuerybyorderidResultDto {
	return poolAlibabaLstMarketingQuerybyorderidResultDto.Get().(*AlibabaLstMarketingQuerybyorderidResultDto)
}

// ReleaseAlibabaLstMarketingQuerybyorderidResultDto 释放AlibabaLstMarketingQuerybyorderidResultDto
func ReleaseAlibabaLstMarketingQuerybyorderidResultDto(v *AlibabaLstMarketingQuerybyorderidResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolAlibabaLstMarketingQuerybyorderidResultDto.Put(v)
}
