package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCrowdAddResultDto 结构体
type TaobaoFeedflowItemCrowdAddResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCrowdAddResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdAddResultDto)
	},
}

// GetTaobaoFeedflowItemCrowdAddResultDto() 从对象池中获取TaobaoFeedflowItemCrowdAddResultDto
func GetTaobaoFeedflowItemCrowdAddResultDto() *TaobaoFeedflowItemCrowdAddResultDto {
	return poolTaobaoFeedflowItemCrowdAddResultDto.Get().(*TaobaoFeedflowItemCrowdAddResultDto)
}

// ReleaseTaobaoFeedflowItemCrowdAddResultDto 释放TaobaoFeedflowItemCrowdAddResultDto
func ReleaseTaobaoFeedflowItemCrowdAddResultDto(v *TaobaoFeedflowItemCrowdAddResultDto) {
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemCrowdAddResultDto.Put(v)
}
