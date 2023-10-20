package tmallgeniescp

import (
	"sync"
)

// IbpChannelDto 结构体
type IbpChannelDto struct {
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 渠道名称
	ChannelName string `json:"channel_name,omitempty" xml:"channel_name,omitempty"`
	// 渠道ID
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
}

var poolIbpChannelDto = sync.Pool{
	New: func() any {
		return new(IbpChannelDto)
	},
}

// GetIbpChannelDto() 从对象池中获取IbpChannelDto
func GetIbpChannelDto() *IbpChannelDto {
	return poolIbpChannelDto.Get().(*IbpChannelDto)
}

// ReleaseIbpChannelDto 释放IbpChannelDto
func ReleaseIbpChannelDto(v *IbpChannelDto) {
	v.Tenant = ""
	v.ChannelName = ""
	v.ChannelId = ""
	v.ExtendJson = ""
	poolIbpChannelDto.Put(v)
}
