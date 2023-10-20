package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupAdzoneBindResultDto 结构体
type TaobaoFeedflowItemAdgroupAdzoneBindResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupAdzoneBindResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupAdzoneBindResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupAdzoneBindResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupAdzoneBindResultDto
func GetTaobaoFeedflowItemAdgroupAdzoneBindResultDto() *TaobaoFeedflowItemAdgroupAdzoneBindResultDto {
	return poolTaobaoFeedflowItemAdgroupAdzoneBindResultDto.Get().(*TaobaoFeedflowItemAdgroupAdzoneBindResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupAdzoneBindResultDto 释放TaobaoFeedflowItemAdgroupAdzoneBindResultDto
func ReleaseTaobaoFeedflowItemAdgroupAdzoneBindResultDto(v *TaobaoFeedflowItemAdgroupAdzoneBindResultDto) {
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemAdgroupAdzoneBindResultDto.Put(v)
}
