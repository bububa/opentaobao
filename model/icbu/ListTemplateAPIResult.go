package icbu

import (
	"sync"
)

// ListTemplateAPIResult 结构体
type ListTemplateAPIResult struct {
	// 运费模板集合
	Items []ShippinglineTemplate `json:"items,omitempty" xml:"items>shippingline_template,omitempty"`
	// 运费模板总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolListTemplateAPIResult = sync.Pool{
	New: func() any {
		return new(ListTemplateAPIResult)
	},
}

// GetListTemplateAPIResult() 从对象池中获取ListTemplateAPIResult
func GetListTemplateAPIResult() *ListTemplateAPIResult {
	return poolListTemplateAPIResult.Get().(*ListTemplateAPIResult)
}

// ReleaseListTemplateAPIResult 释放ListTemplateAPIResult
func ReleaseListTemplateAPIResult(v *ListTemplateAPIResult) {
	v.Items = v.Items[:0]
	v.Total = 0
	poolListTemplateAPIResult.Put(v)
}
