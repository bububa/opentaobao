package trade

import (
	"sync"
)

// ChannelOrderOption 结构体
type ChannelOrderOption struct {
	// 是否允许供应商修改
	IsAllowUpperModify bool `json:"is_allow_upper_modify,omitempty" xml:"is_allow_upper_modify,omitempty"`
}

var poolChannelOrderOption = sync.Pool{
	New: func() any {
		return new(ChannelOrderOption)
	},
}

// GetChannelOrderOption() 从对象池中获取ChannelOrderOption
func GetChannelOrderOption() *ChannelOrderOption {
	return poolChannelOrderOption.Get().(*ChannelOrderOption)
}

// ReleaseChannelOrderOption 释放ChannelOrderOption
func ReleaseChannelOrderOption(v *ChannelOrderOption) {
	v.IsAllowUpperModify = false
	poolChannelOrderOption.Put(v)
}
