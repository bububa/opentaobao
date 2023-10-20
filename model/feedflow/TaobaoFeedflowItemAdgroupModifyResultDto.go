package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupModifyResultDto 结构体
type TaobaoFeedflowItemAdgroupModifyResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 单元id
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupModifyResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupModifyResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupModifyResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupModifyResultDto
func GetTaobaoFeedflowItemAdgroupModifyResultDto() *TaobaoFeedflowItemAdgroupModifyResultDto {
	return poolTaobaoFeedflowItemAdgroupModifyResultDto.Get().(*TaobaoFeedflowItemAdgroupModifyResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupModifyResultDto 释放TaobaoFeedflowItemAdgroupModifyResultDto
func ReleaseTaobaoFeedflowItemAdgroupModifyResultDto(v *TaobaoFeedflowItemAdgroupModifyResultDto) {
	v.Message = ""
	v.Result = false
	v.Success = false
	poolTaobaoFeedflowItemAdgroupModifyResultDto.Put(v)
}
