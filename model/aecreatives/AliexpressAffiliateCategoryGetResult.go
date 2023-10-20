package aecreatives

import (
	"sync"
)

// AliexpressAffiliateCategoryGetResult 结构体
type AliexpressAffiliateCategoryGetResult struct {
	// 类目信息
	Categories []Category `json:"categories,omitempty" xml:"categories>category,omitempty"`
	// 返回结果数量
	TotalResultCount int64 `json:"total_result_count,omitempty" xml:"total_result_count,omitempty"`
}

var poolAliexpressAffiliateCategoryGetResult = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateCategoryGetResult)
	},
}

// GetAliexpressAffiliateCategoryGetResult() 从对象池中获取AliexpressAffiliateCategoryGetResult
func GetAliexpressAffiliateCategoryGetResult() *AliexpressAffiliateCategoryGetResult {
	return poolAliexpressAffiliateCategoryGetResult.Get().(*AliexpressAffiliateCategoryGetResult)
}

// ReleaseAliexpressAffiliateCategoryGetResult 释放AliexpressAffiliateCategoryGetResult
func ReleaseAliexpressAffiliateCategoryGetResult(v *AliexpressAffiliateCategoryGetResult) {
	v.Categories = v.Categories[:0]
	v.TotalResultCount = 0
	poolAliexpressAffiliateCategoryGetResult.Put(v)
}
