package legalsuit

import (
	"sync"
)

// SceneOption 结构体
type SceneOption struct {
	// 场景名称
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 场景code
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolSceneOption = sync.Pool{
	New: func() any {
		return new(SceneOption)
	},
}

// GetSceneOption() 从对象池中获取SceneOption
func GetSceneOption() *SceneOption {
	return poolSceneOption.Get().(*SceneOption)
}

// ReleaseSceneOption 释放SceneOption
func ReleaseSceneOption(v *SceneOption) {
	v.Text = ""
	v.Value = ""
	poolSceneOption.Put(v)
}
