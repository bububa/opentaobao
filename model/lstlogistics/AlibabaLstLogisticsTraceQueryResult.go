package lstlogistics

import (
	"sync"
)

// AlibabaLstLogisticsTraceQueryResult 结构体
type AlibabaLstLogisticsTraceQueryResult struct {
	// 返回内容
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstLogisticsTraceQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstLogisticsTraceQueryResult)
	},
}

// GetAlibabaLstLogisticsTraceQueryResult() 从对象池中获取AlibabaLstLogisticsTraceQueryResult
func GetAlibabaLstLogisticsTraceQueryResult() *AlibabaLstLogisticsTraceQueryResult {
	return poolAlibabaLstLogisticsTraceQueryResult.Get().(*AlibabaLstLogisticsTraceQueryResult)
}

// ReleaseAlibabaLstLogisticsTraceQueryResult 释放AlibabaLstLogisticsTraceQueryResult
func ReleaseAlibabaLstLogisticsTraceQueryResult(v *AlibabaLstLogisticsTraceQueryResult) {
	v.ContentList = v.ContentList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaLstLogisticsTraceQueryResult.Put(v)
}
