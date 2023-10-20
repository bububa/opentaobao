package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto 结构体
type TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupAdzoneUnbindResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupAdzoneUnbindResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto
func GetTaobaoFeedflowItemAdgroupAdzoneUnbindResultDto() *TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto {
	return poolTaobaoFeedflowItemAdgroupAdzoneUnbindResultDto.Get().(*TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupAdzoneUnbindResultDto 释放TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto
func ReleaseTaobaoFeedflowItemAdgroupAdzoneUnbindResultDto(v *TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto) {
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemAdgroupAdzoneUnbindResultDto.Put(v)
}
