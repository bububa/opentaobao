package idleparttime

import (
	"sync"
)

// PartTimeRequireSchema 结构体
type PartTimeRequireSchema struct {
	// 要求
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 要求描述
	Requirement string `json:"requirement,omitempty" xml:"requirement,omitempty"`
	// 类型, 1:文本 2: 图片
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolPartTimeRequireSchema = sync.Pool{
	New: func() any {
		return new(PartTimeRequireSchema)
	},
}

// GetPartTimeRequireSchema() 从对象池中获取PartTimeRequireSchema
func GetPartTimeRequireSchema() *PartTimeRequireSchema {
	return poolPartTimeRequireSchema.Get().(*PartTimeRequireSchema)
}

// ReleasePartTimeRequireSchema 释放PartTimeRequireSchema
func ReleasePartTimeRequireSchema(v *PartTimeRequireSchema) {
	v.Description = ""
	v.Requirement = ""
	v.Type = 0
	poolPartTimeRequireSchema.Put(v)
}
