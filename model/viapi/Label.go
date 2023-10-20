package viapi

import (
	"sync"
)

// Label 结构体
type Label struct {
	// 可选值包括：  spam：含垃圾信息 politics： 涉政 abuse：辱骂 porn：智能鉴黄 terrorism：暴恐识别 flood：灌水 contraband：违禁 ad：文本违规识别
	Label string `json:"label,omitempty" xml:"label,omitempty"`
}

var poolLabel = sync.Pool{
	New: func() any {
		return new(Label)
	},
}

// GetLabel() 从对象池中获取Label
func GetLabel() *Label {
	return poolLabel.Get().(*Label)
}

// ReleaseLabel 释放Label
func ReleaseLabel(v *Label) {
	v.Label = ""
	poolLabel.Put(v)
}
