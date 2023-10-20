package campus

import (
	"sync"
)

// TemplateApiQuery 结构体
type TemplateApiQuery struct {
	// 设备模板编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 模糊查询模板编码或名称
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 设备模板名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 分页大小,限制limit&lt;=500
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolTemplateApiQuery = sync.Pool{
	New: func() any {
		return new(TemplateApiQuery)
	},
}

// GetTemplateApiQuery() 从对象池中获取TemplateApiQuery
func GetTemplateApiQuery() *TemplateApiQuery {
	return poolTemplateApiQuery.Get().(*TemplateApiQuery)
}

// ReleaseTemplateApiQuery 释放TemplateApiQuery
func ReleaseTemplateApiQuery(v *TemplateApiQuery) {
	v.Code = ""
	v.Key = ""
	v.Name = ""
	v.Limit = 0
	v.CurrentPage = 0
	poolTemplateApiQuery.Put(v)
}
