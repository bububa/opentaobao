package alsc

import (
	"sync"
)

// ActionOpenInfo 结构体
type ActionOpenInfo struct {
	// 行动点描述
	ActionDesc string `json:"action_desc,omitempty" xml:"action_desc,omitempty"`
	// 行动点类型
	ActionType string `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 跳转链接
	JumpUrl string `json:"jump_url,omitempty" xml:"jump_url,omitempty"`
	// 文案是否高亮
	Highlight bool `json:"highlight,omitempty" xml:"highlight,omitempty"`
}

var poolActionOpenInfo = sync.Pool{
	New: func() any {
		return new(ActionOpenInfo)
	},
}

// GetActionOpenInfo() 从对象池中获取ActionOpenInfo
func GetActionOpenInfo() *ActionOpenInfo {
	return poolActionOpenInfo.Get().(*ActionOpenInfo)
}

// ReleaseActionOpenInfo 释放ActionOpenInfo
func ReleaseActionOpenInfo(v *ActionOpenInfo) {
	v.ActionDesc = ""
	v.ActionType = ""
	v.JumpUrl = ""
	v.Highlight = false
	poolActionOpenInfo.Put(v)
}
