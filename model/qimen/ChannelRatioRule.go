package qimen

import (
	"sync"
)

// ChannelRatioRule 结构体
type ChannelRatioRule struct {
	// 奇门仓储字段,C1223,string(50),,
	ChannelCode string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	// 奇门仓储字段,C1223,string(50),,
	Ratio string `json:"ratio,omitempty" xml:"ratio,omitempty"`
}

var poolChannelRatioRule = sync.Pool{
	New: func() any {
		return new(ChannelRatioRule)
	},
}

// GetChannelRatioRule() 从对象池中获取ChannelRatioRule
func GetChannelRatioRule() *ChannelRatioRule {
	return poolChannelRatioRule.Get().(*ChannelRatioRule)
}

// ReleaseChannelRatioRule 释放ChannelRatioRule
func ReleaseChannelRatioRule(v *ChannelRatioRule) {
	v.ChannelCode = ""
	v.Ratio = ""
	poolChannelRatioRule.Put(v)
}
