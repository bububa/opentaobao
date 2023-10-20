package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemOptionPageResultDto 结构体
type TaobaoFeedflowItemOptionPageResultDto struct {
	// 标签信息
	Labels []LabelDto `json:"labels,omitempty" xml:"labels>label_dto,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemOptionPageResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemOptionPageResultDto)
	},
}

// GetTaobaoFeedflowItemOptionPageResultDto() 从对象池中获取TaobaoFeedflowItemOptionPageResultDto
func GetTaobaoFeedflowItemOptionPageResultDto() *TaobaoFeedflowItemOptionPageResultDto {
	return poolTaobaoFeedflowItemOptionPageResultDto.Get().(*TaobaoFeedflowItemOptionPageResultDto)
}

// ReleaseTaobaoFeedflowItemOptionPageResultDto 释放TaobaoFeedflowItemOptionPageResultDto
func ReleaseTaobaoFeedflowItemOptionPageResultDto(v *TaobaoFeedflowItemOptionPageResultDto) {
	v.Labels = v.Labels[:0]
	v.Message = ""
	v.TotalCount = 0
	v.Success = false
	poolTaobaoFeedflowItemOptionPageResultDto.Put(v)
}
