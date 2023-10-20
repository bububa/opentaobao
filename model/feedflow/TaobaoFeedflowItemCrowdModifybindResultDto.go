package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCrowdModifybindResultDto 结构体
type TaobaoFeedflowItemCrowdModifybindResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCrowdModifybindResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdModifybindResultDto)
	},
}

// GetTaobaoFeedflowItemCrowdModifybindResultDto() 从对象池中获取TaobaoFeedflowItemCrowdModifybindResultDto
func GetTaobaoFeedflowItemCrowdModifybindResultDto() *TaobaoFeedflowItemCrowdModifybindResultDto {
	return poolTaobaoFeedflowItemCrowdModifybindResultDto.Get().(*TaobaoFeedflowItemCrowdModifybindResultDto)
}

// ReleaseTaobaoFeedflowItemCrowdModifybindResultDto 释放TaobaoFeedflowItemCrowdModifybindResultDto
func ReleaseTaobaoFeedflowItemCrowdModifybindResultDto(v *TaobaoFeedflowItemCrowdModifybindResultDto) {
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemCrowdModifybindResultDto.Put(v)
}
