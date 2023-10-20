package kclub

import (
	"sync"
)

// AlibabaKclubKcQueryknowledgeResult 结构体
type AlibabaKclubKcQueryknowledgeResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回数据
	Data *Paging `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaKclubKcQueryknowledgeResult = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcQueryknowledgeResult)
	},
}

// GetAlibabaKclubKcQueryknowledgeResult() 从对象池中获取AlibabaKclubKcQueryknowledgeResult
func GetAlibabaKclubKcQueryknowledgeResult() *AlibabaKclubKcQueryknowledgeResult {
	return poolAlibabaKclubKcQueryknowledgeResult.Get().(*AlibabaKclubKcQueryknowledgeResult)
}

// ReleaseAlibabaKclubKcQueryknowledgeResult 释放AlibabaKclubKcQueryknowledgeResult
func ReleaseAlibabaKclubKcQueryknowledgeResult(v *AlibabaKclubKcQueryknowledgeResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaKclubKcQueryknowledgeResult.Put(v)
}
