package security

import (
	"sync"
)

// PluginSummary 结构体
type PluginSummary struct {
	// 插件个数
	PluginCount int64 `json:"plugin_count,omitempty" xml:"plugin_count,omitempty"`
	// 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolPluginSummary = sync.Pool{
	New: func() any {
		return new(PluginSummary)
	},
}

// GetPluginSummary() 从对象池中获取PluginSummary
func GetPluginSummary() *PluginSummary {
	return poolPluginSummary.Get().(*PluginSummary)
}

// ReleasePluginSummary 释放PluginSummary
func ReleasePluginSummary(v *PluginSummary) {
	v.PluginCount = 0
	v.Status = 0
	poolPluginSummary.Put(v)
}
