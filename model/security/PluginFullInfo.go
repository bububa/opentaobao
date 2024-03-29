package security

import (
	"sync"
)

// PluginFullInfo 结构体
type PluginFullInfo struct {
	// 插件详细信息
	PluginDetails []PluginDetail `json:"plugin_details,omitempty" xml:"plugin_details>plugin_detail,omitempty"`
	// 插件个数
	PluginCount int64 `json:"plugin_count,omitempty" xml:"plugin_count,omitempty"`
	// 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolPluginFullInfo = sync.Pool{
	New: func() any {
		return new(PluginFullInfo)
	},
}

// GetPluginFullInfo() 从对象池中获取PluginFullInfo
func GetPluginFullInfo() *PluginFullInfo {
	return poolPluginFullInfo.Get().(*PluginFullInfo)
}

// ReleasePluginFullInfo 释放PluginFullInfo
func ReleasePluginFullInfo(v *PluginFullInfo) {
	v.PluginDetails = v.PluginDetails[:0]
	v.PluginCount = 0
	v.Status = 0
	poolPluginFullInfo.Put(v)
}
