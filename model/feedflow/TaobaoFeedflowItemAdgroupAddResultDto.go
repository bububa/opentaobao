package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupAddResultDto 结构体
type TaobaoFeedflowItemAdgroupAddResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 单元id
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupAddResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupAddResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupAddResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupAddResultDto
func GetTaobaoFeedflowItemAdgroupAddResultDto() *TaobaoFeedflowItemAdgroupAddResultDto {
	return poolTaobaoFeedflowItemAdgroupAddResultDto.Get().(*TaobaoFeedflowItemAdgroupAddResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupAddResultDto 释放TaobaoFeedflowItemAdgroupAddResultDto
func ReleaseTaobaoFeedflowItemAdgroupAddResultDto(v *TaobaoFeedflowItemAdgroupAddResultDto) {
	v.Message = ""
	v.Result = 0
	v.Success = false
	poolTaobaoFeedflowItemAdgroupAddResultDto.Put(v)
}
