package dmp

import (
	"sync"
)

// TopicQueryDto 结构体
type TopicQueryDto struct {
	// 算法推荐模版context
	ContextParams string `json:"context_params,omitempty" xml:"context_params,omitempty"`
}

var poolTopicQueryDto = sync.Pool{
	New: func() any {
		return new(TopicQueryDto)
	},
}

// GetTopicQueryDto() 从对象池中获取TopicQueryDto
func GetTopicQueryDto() *TopicQueryDto {
	return poolTopicQueryDto.Get().(*TopicQueryDto)
}

// ReleaseTopicQueryDto 释放TopicQueryDto
func ReleaseTopicQueryDto(v *TopicQueryDto) {
	v.ContextParams = ""
	poolTopicQueryDto.Put(v)
}
