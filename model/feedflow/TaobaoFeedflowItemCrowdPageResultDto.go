package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCrowdPageResultDto 结构体
type TaobaoFeedflowItemCrowdPageResultDto struct {
	// 人群列表
	Crowds []CrowdDto `json:"crowds,omitempty" xml:"crowds>crowd_dto,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCrowdPageResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdPageResultDto)
	},
}

// GetTaobaoFeedflowItemCrowdPageResultDto() 从对象池中获取TaobaoFeedflowItemCrowdPageResultDto
func GetTaobaoFeedflowItemCrowdPageResultDto() *TaobaoFeedflowItemCrowdPageResultDto {
	return poolTaobaoFeedflowItemCrowdPageResultDto.Get().(*TaobaoFeedflowItemCrowdPageResultDto)
}

// ReleaseTaobaoFeedflowItemCrowdPageResultDto 释放TaobaoFeedflowItemCrowdPageResultDto
func ReleaseTaobaoFeedflowItemCrowdPageResultDto(v *TaobaoFeedflowItemCrowdPageResultDto) {
	v.Crowds = v.Crowds[:0]
	v.Message = ""
	v.TotalCount = 0
	v.Success = false
	poolTaobaoFeedflowItemCrowdPageResultDto.Put(v)
}
