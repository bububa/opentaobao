package security

import (
	"sync"
)

// ChannelAppInfo 结构体
type ChannelAppInfo struct {
	// 渠道名称,多渠道加固才有值
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 加固后的APP下载地址
	AppUrl string `json:"app_url,omitempty" xml:"app_url,omitempty"`
}

var poolChannelAppInfo = sync.Pool{
	New: func() any {
		return new(ChannelAppInfo)
	},
}

// GetChannelAppInfo() 从对象池中获取ChannelAppInfo
func GetChannelAppInfo() *ChannelAppInfo {
	return poolChannelAppInfo.Get().(*ChannelAppInfo)
}

// ReleaseChannelAppInfo 释放ChannelAppInfo
func ReleaseChannelAppInfo(v *ChannelAppInfo) {
	v.Channel = ""
	v.AppUrl = ""
	poolChannelAppInfo.Put(v)
}
