package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupDeleteResultDto 结构体
type TaobaoFeedflowItemAdgroupDeleteResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 删除结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupDeleteResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupDeleteResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupDeleteResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupDeleteResultDto
func GetTaobaoFeedflowItemAdgroupDeleteResultDto() *TaobaoFeedflowItemAdgroupDeleteResultDto {
	return poolTaobaoFeedflowItemAdgroupDeleteResultDto.Get().(*TaobaoFeedflowItemAdgroupDeleteResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupDeleteResultDto 释放TaobaoFeedflowItemAdgroupDeleteResultDto
func ReleaseTaobaoFeedflowItemAdgroupDeleteResultDto(v *TaobaoFeedflowItemAdgroupDeleteResultDto) {
	v.Message = ""
	v.Result = false
	v.Success = false
	poolTaobaoFeedflowItemAdgroupDeleteResultDto.Put(v)
}
