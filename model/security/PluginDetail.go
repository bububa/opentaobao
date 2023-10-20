package security

import (
	"sync"
)

// PluginDetail 结构体
type PluginDetail struct {
	// 插件行为
	Actions []string `json:"actions,omitempty" xml:"actions>string,omitempty"`
	// 插件类型
	Types []string `json:"types,omitempty" xml:"types>string,omitempty"`
	// 插件开发商
	Company string `json:"company,omitempty" xml:"company,omitempty"`
	// 插件描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 插件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 插件位置
	Path string `json:"path,omitempty" xml:"path,omitempty"`
}

var poolPluginDetail = sync.Pool{
	New: func() any {
		return new(PluginDetail)
	},
}

// GetPluginDetail() 从对象池中获取PluginDetail
func GetPluginDetail() *PluginDetail {
	return poolPluginDetail.Get().(*PluginDetail)
}

// ReleasePluginDetail 释放PluginDetail
func ReleasePluginDetail(v *PluginDetail) {
	v.Actions = v.Actions[:0]
	v.Types = v.Types[:0]
	v.Company = ""
	v.Desc = ""
	v.Name = ""
	v.Path = ""
	poolPluginDetail.Put(v)
}
