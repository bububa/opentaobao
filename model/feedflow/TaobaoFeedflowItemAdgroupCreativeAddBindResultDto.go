package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupCreativeAddBindResultDto 结构体
type TaobaoFeedflowItemAdgroupCreativeAddBindResultDto struct {
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupCreativeAddBindResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupCreativeAddBindResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupCreativeAddBindResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupCreativeAddBindResultDto
func GetTaobaoFeedflowItemAdgroupCreativeAddBindResultDto() *TaobaoFeedflowItemAdgroupCreativeAddBindResultDto {
	return poolTaobaoFeedflowItemAdgroupCreativeAddBindResultDto.Get().(*TaobaoFeedflowItemAdgroupCreativeAddBindResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupCreativeAddBindResultDto 释放TaobaoFeedflowItemAdgroupCreativeAddBindResultDto
func ReleaseTaobaoFeedflowItemAdgroupCreativeAddBindResultDto(v *TaobaoFeedflowItemAdgroupCreativeAddBindResultDto) {
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemAdgroupCreativeAddBindResultDto.Put(v)
}
