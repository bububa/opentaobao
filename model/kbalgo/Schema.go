package kbalgo

import (
	"sync"
)

// Schema 结构体
type Schema struct {
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 页面类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// lbs信息
	Lbs string `json:"lbs,omitempty" xml:"lbs,omitempty"`
}

var poolSchema = sync.Pool{
	New: func() any {
		return new(Schema)
	},
}

// GetSchema() 从对象池中获取Schema
func GetSchema() *Schema {
	return poolSchema.Get().(*Schema)
}

// ReleaseSchema 释放Schema
func ReleaseSchema(v *Schema) {
	v.Url = ""
	v.Type = ""
	v.Lbs = ""
	poolSchema.Put(v)
}
