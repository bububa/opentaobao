package alitripbp

import (
	"sync"
)

// ChannelExamineResultDto 结构体
type ChannelExamineResultDto struct {
	// 活动url
	TargetUrl string `json:"target_url,omitempty" xml:"target_url,omitempty"`
	// 是否为活动的目标用户
	IsTargetCrow bool `json:"is_target_crow,omitempty" xml:"is_target_crow,omitempty"`
}

var poolChannelExamineResultDto = sync.Pool{
	New: func() any {
		return new(ChannelExamineResultDto)
	},
}

// GetChannelExamineResultDto() 从对象池中获取ChannelExamineResultDto
func GetChannelExamineResultDto() *ChannelExamineResultDto {
	return poolChannelExamineResultDto.Get().(*ChannelExamineResultDto)
}

// ReleaseChannelExamineResultDto 释放ChannelExamineResultDto
func ReleaseChannelExamineResultDto(v *ChannelExamineResultDto) {
	v.TargetUrl = ""
	v.IsTargetCrow = false
	poolChannelExamineResultDto.Put(v)
}
