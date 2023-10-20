package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCrowdModifyResultDto 结构体
type TaobaoFeedflowItemCrowdModifyResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCrowdModifyResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdModifyResultDto)
	},
}

// GetTaobaoFeedflowItemCrowdModifyResultDto() 从对象池中获取TaobaoFeedflowItemCrowdModifyResultDto
func GetTaobaoFeedflowItemCrowdModifyResultDto() *TaobaoFeedflowItemCrowdModifyResultDto {
	return poolTaobaoFeedflowItemCrowdModifyResultDto.Get().(*TaobaoFeedflowItemCrowdModifyResultDto)
}

// ReleaseTaobaoFeedflowItemCrowdModifyResultDto 释放TaobaoFeedflowItemCrowdModifyResultDto
func ReleaseTaobaoFeedflowItemCrowdModifyResultDto(v *TaobaoFeedflowItemCrowdModifyResultDto) {
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemCrowdModifyResultDto.Put(v)
}
