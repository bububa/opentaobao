package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupPageResultDto 结构体
type TaobaoFeedflowItemAdgroupPageResultDto struct {
	// 返回数据结果
	Results []AdgroupDto `json:"results,omitempty" xml:"results>adgroup_dto,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupPageResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupPageResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupPageResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupPageResultDto
func GetTaobaoFeedflowItemAdgroupPageResultDto() *TaobaoFeedflowItemAdgroupPageResultDto {
	return poolTaobaoFeedflowItemAdgroupPageResultDto.Get().(*TaobaoFeedflowItemAdgroupPageResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupPageResultDto 释放TaobaoFeedflowItemAdgroupPageResultDto
func ReleaseTaobaoFeedflowItemAdgroupPageResultDto(v *TaobaoFeedflowItemAdgroupPageResultDto) {
	v.Results = v.Results[:0]
	v.Message = ""
	v.TotalCount = 0
	v.Success = false
	poolTaobaoFeedflowItemAdgroupPageResultDto.Put(v)
}
