package kclub

import (
	"sync"
)

// SorterConfig 结构体
type SorterConfig struct {
	// 排序顺序
	Order string `json:"order,omitempty" xml:"order,omitempty"`
	// 排序字段
	Field string `json:"field,omitempty" xml:"field,omitempty"`
}

var poolSorterConfig = sync.Pool{
	New: func() any {
		return new(SorterConfig)
	},
}

// GetSorterConfig() 从对象池中获取SorterConfig
func GetSorterConfig() *SorterConfig {
	return poolSorterConfig.Get().(*SorterConfig)
}

// ReleaseSorterConfig 释放SorterConfig
func ReleaseSorterConfig(v *SorterConfig) {
	v.Order = ""
	v.Field = ""
	poolSorterConfig.Put(v)
}
