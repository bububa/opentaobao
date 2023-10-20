package wdk

import (
	"sync"
)

// ChannelProp 结构体
type ChannelProp struct {
	// 渠道属性，取值为key-value键值对形式
	Props []PropField `json:"props,omitempty" xml:"props>prop_field,omitempty"`
	// 渠道类型：txd淘鲜达，elm饿了么，shareStore共享库存
	ChannelType string `json:"channel_type,omitempty" xml:"channel_type,omitempty"`
}

var poolChannelProp = sync.Pool{
	New: func() any {
		return new(ChannelProp)
	},
}

// GetChannelProp() 从对象池中获取ChannelProp
func GetChannelProp() *ChannelProp {
	return poolChannelProp.Get().(*ChannelProp)
}

// ReleaseChannelProp 释放ChannelProp
func ReleaseChannelProp(v *ChannelProp) {
	v.Props = v.Props[:0]
	v.ChannelType = ""
	poolChannelProp.Put(v)
}
