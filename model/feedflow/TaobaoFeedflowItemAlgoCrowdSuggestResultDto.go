package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAlgoCrowdSuggestResultDto 结构体
type TaobaoFeedflowItemAlgoCrowdSuggestResultDto struct {
	// 人群列表
	Crowds []CrowdDto `json:"crowds,omitempty" xml:"crowds>crowd_dto,omitempty"`
	// 失败时候的消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAlgoCrowdSuggestResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAlgoCrowdSuggestResultDto)
	},
}

// GetTaobaoFeedflowItemAlgoCrowdSuggestResultDto() 从对象池中获取TaobaoFeedflowItemAlgoCrowdSuggestResultDto
func GetTaobaoFeedflowItemAlgoCrowdSuggestResultDto() *TaobaoFeedflowItemAlgoCrowdSuggestResultDto {
	return poolTaobaoFeedflowItemAlgoCrowdSuggestResultDto.Get().(*TaobaoFeedflowItemAlgoCrowdSuggestResultDto)
}

// ReleaseTaobaoFeedflowItemAlgoCrowdSuggestResultDto 释放TaobaoFeedflowItemAlgoCrowdSuggestResultDto
func ReleaseTaobaoFeedflowItemAlgoCrowdSuggestResultDto(v *TaobaoFeedflowItemAlgoCrowdSuggestResultDto) {
	v.Crowds = v.Crowds[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemAlgoCrowdSuggestResultDto.Put(v)
}
