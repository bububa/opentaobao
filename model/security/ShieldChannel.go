package security

import (
	"sync"
)

// ShieldChannel 结构体
type ShieldChannel struct {
	// 渠道名称列表
	Values []string `json:"values,omitempty" xml:"values>string,omitempty"`
	// AndroidManifest.xml 中用于表示渠道信息的meta-data标签的android:name
	MetaName string `json:"meta_name,omitempty" xml:"meta_name,omitempty"`
}

var poolShieldChannel = sync.Pool{
	New: func() any {
		return new(ShieldChannel)
	},
}

// GetShieldChannel() 从对象池中获取ShieldChannel
func GetShieldChannel() *ShieldChannel {
	return poolShieldChannel.Get().(*ShieldChannel)
}

// ReleaseShieldChannel 释放ShieldChannel
func ReleaseShieldChannel(v *ShieldChannel) {
	v.Values = v.Values[:0]
	v.MetaName = ""
	poolShieldChannel.Put(v)
}
