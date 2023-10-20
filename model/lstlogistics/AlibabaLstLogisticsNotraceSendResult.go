package lstlogistics

import (
	"sync"
)

// AlibabaLstLogisticsNotraceSendResult 结构体
type AlibabaLstLogisticsNotraceSendResult struct {
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstLogisticsNotraceSendResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstLogisticsNotraceSendResult)
	},
}

// GetAlibabaLstLogisticsNotraceSendResult() 从对象池中获取AlibabaLstLogisticsNotraceSendResult
func GetAlibabaLstLogisticsNotraceSendResult() *AlibabaLstLogisticsNotraceSendResult {
	return poolAlibabaLstLogisticsNotraceSendResult.Get().(*AlibabaLstLogisticsNotraceSendResult)
}

// ReleaseAlibabaLstLogisticsNotraceSendResult 释放AlibabaLstLogisticsNotraceSendResult
func ReleaseAlibabaLstLogisticsNotraceSendResult(v *AlibabaLstLogisticsNotraceSendResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolAlibabaLstLogisticsNotraceSendResult.Put(v)
}
