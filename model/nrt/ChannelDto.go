package nrt

import (
	"sync"
)

// ChannelDto 结构体
type ChannelDto struct {
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 渠道名称
	ChannelName string `json:"channel_name,omitempty" xml:"channel_name,omitempty"`
}

var poolChannelDto = sync.Pool{
	New: func() any {
		return new(ChannelDto)
	},
}

// GetChannelDto() 从对象池中获取ChannelDto
func GetChannelDto() *ChannelDto {
	return poolChannelDto.Get().(*ChannelDto)
}

// ReleaseChannelDto 释放ChannelDto
func ReleaseChannelDto(v *ChannelDto) {
	v.ChannelCode = ""
	v.ChannelName = ""
	poolChannelDto.Put(v)
}
