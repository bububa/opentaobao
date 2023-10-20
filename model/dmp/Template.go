package dmp

import (
	"sync"
)

// Template 结构体
type Template struct {
	// 模版扩展json信息，usageIndex: 使用热度，mark: 打标文案
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 模版描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 模版推荐理由
	Highlight string `json:"highlight,omitempty" xml:"highlight,omitempty"`
	// 模版名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 模版id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTemplate = sync.Pool{
	New: func() any {
		return new(Template)
	},
}

// GetTemplate() 从对象池中获取Template
func GetTemplate() *Template {
	return poolTemplate.Get().(*Template)
}

// ReleaseTemplate 释放Template
func ReleaseTemplate(v *Template) {
	v.Ext = ""
	v.Description = ""
	v.Highlight = ""
	v.Name = ""
	v.Id = 0
	poolTemplate.Put(v)
}
