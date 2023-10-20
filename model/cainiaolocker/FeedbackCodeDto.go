package cainiaolocker

import (
	"sync"
)

// FeedbackCodeDto 结构体
type FeedbackCodeDto struct {
	// 异常反馈编码
	FeedbackCode string `json:"feedback_code,omitempty" xml:"feedback_code,omitempty"`
	// 异常反馈编码描述
	FeedbackDesc string `json:"feedback_desc,omitempty" xml:"feedback_desc,omitempty"`
}

var poolFeedbackCodeDto = sync.Pool{
	New: func() any {
		return new(FeedbackCodeDto)
	},
}

// GetFeedbackCodeDto() 从对象池中获取FeedbackCodeDto
func GetFeedbackCodeDto() *FeedbackCodeDto {
	return poolFeedbackCodeDto.Get().(*FeedbackCodeDto)
}

// ReleaseFeedbackCodeDto 释放FeedbackCodeDto
func ReleaseFeedbackCodeDto(v *FeedbackCodeDto) {
	v.FeedbackCode = ""
	v.FeedbackDesc = ""
	poolFeedbackCodeDto.Put(v)
}
