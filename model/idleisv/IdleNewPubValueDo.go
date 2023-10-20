package idleisv

import (
	"sync"
)

// IdleNewPubValueDo 结构体
type IdleNewPubValueDo struct {
	// 属性id
	PropertyId string `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 属性名称
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
	// 渠道类目id
	ChannelCatId string `json:"channel_cat_id,omitempty" xml:"channel_cat_id,omitempty"`
	// 值id
	ValueId string `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 值名称
	ValueName string `json:"value_name,omitempty" xml:"value_name,omitempty"`
}

var poolIdleNewPubValueDo = sync.Pool{
	New: func() any {
		return new(IdleNewPubValueDo)
	},
}

// GetIdleNewPubValueDo() 从对象池中获取IdleNewPubValueDo
func GetIdleNewPubValueDo() *IdleNewPubValueDo {
	return poolIdleNewPubValueDo.Get().(*IdleNewPubValueDo)
}

// ReleaseIdleNewPubValueDo 释放IdleNewPubValueDo
func ReleaseIdleNewPubValueDo(v *IdleNewPubValueDo) {
	v.PropertyId = ""
	v.PropertyName = ""
	v.ChannelCatId = ""
	v.ValueId = ""
	v.ValueName = ""
	poolIdleNewPubValueDo.Put(v)
}
