package lstlogistics

import (
	"sync"
)

// AlibabaLstLogisticsThirdpartSendResult 结构体
type AlibabaLstLogisticsThirdpartSendResult struct {
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstLogisticsThirdpartSendResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstLogisticsThirdpartSendResult)
	},
}

// GetAlibabaLstLogisticsThirdpartSendResult() 从对象池中获取AlibabaLstLogisticsThirdpartSendResult
func GetAlibabaLstLogisticsThirdpartSendResult() *AlibabaLstLogisticsThirdpartSendResult {
	return poolAlibabaLstLogisticsThirdpartSendResult.Get().(*AlibabaLstLogisticsThirdpartSendResult)
}

// ReleaseAlibabaLstLogisticsThirdpartSendResult 释放AlibabaLstLogisticsThirdpartSendResult
func ReleaseAlibabaLstLogisticsThirdpartSendResult(v *AlibabaLstLogisticsThirdpartSendResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolAlibabaLstLogisticsThirdpartSendResult.Put(v)
}
