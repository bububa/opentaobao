package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemTargetValidlistResultDto 结构体
type TaobaoFeedflowItemTargetValidlistResultDto struct {
	// 定向结构
	Targets []TargetDto `json:"targets,omitempty" xml:"targets>target_dto,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemTargetValidlistResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemTargetValidlistResultDto)
	},
}

// GetTaobaoFeedflowItemTargetValidlistResultDto() 从对象池中获取TaobaoFeedflowItemTargetValidlistResultDto
func GetTaobaoFeedflowItemTargetValidlistResultDto() *TaobaoFeedflowItemTargetValidlistResultDto {
	return poolTaobaoFeedflowItemTargetValidlistResultDto.Get().(*TaobaoFeedflowItemTargetValidlistResultDto)
}

// ReleaseTaobaoFeedflowItemTargetValidlistResultDto 释放TaobaoFeedflowItemTargetValidlistResultDto
func ReleaseTaobaoFeedflowItemTargetValidlistResultDto(v *TaobaoFeedflowItemTargetValidlistResultDto) {
	v.Targets = v.Targets[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemTargetValidlistResultDto.Put(v)
}
