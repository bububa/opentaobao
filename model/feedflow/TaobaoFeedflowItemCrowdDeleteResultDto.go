package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCrowdDeleteResultDto 结构体
type TaobaoFeedflowItemCrowdDeleteResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCrowdDeleteResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdDeleteResultDto)
	},
}

// GetTaobaoFeedflowItemCrowdDeleteResultDto() 从对象池中获取TaobaoFeedflowItemCrowdDeleteResultDto
func GetTaobaoFeedflowItemCrowdDeleteResultDto() *TaobaoFeedflowItemCrowdDeleteResultDto {
	return poolTaobaoFeedflowItemCrowdDeleteResultDto.Get().(*TaobaoFeedflowItemCrowdDeleteResultDto)
}

// ReleaseTaobaoFeedflowItemCrowdDeleteResultDto 释放TaobaoFeedflowItemCrowdDeleteResultDto
func ReleaseTaobaoFeedflowItemCrowdDeleteResultDto(v *TaobaoFeedflowItemCrowdDeleteResultDto) {
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemCrowdDeleteResultDto.Put(v)
}
