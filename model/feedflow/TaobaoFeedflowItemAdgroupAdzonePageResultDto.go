package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupAdzonePageResultDto 结构体
type TaobaoFeedflowItemAdgroupAdzonePageResultDto struct {
	// 广告位列表
	AdzoneBindList []AdzoneBindDto `json:"adzone_bind_list,omitempty" xml:"adzone_bind_list>adzone_bind_dto,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 广告位总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupAdzonePageResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupAdzonePageResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupAdzonePageResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupAdzonePageResultDto
func GetTaobaoFeedflowItemAdgroupAdzonePageResultDto() *TaobaoFeedflowItemAdgroupAdzonePageResultDto {
	return poolTaobaoFeedflowItemAdgroupAdzonePageResultDto.Get().(*TaobaoFeedflowItemAdgroupAdzonePageResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupAdzonePageResultDto 释放TaobaoFeedflowItemAdgroupAdzonePageResultDto
func ReleaseTaobaoFeedflowItemAdgroupAdzonePageResultDto(v *TaobaoFeedflowItemAdgroupAdzonePageResultDto) {
	v.AdzoneBindList = v.AdzoneBindList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.Success = false
	poolTaobaoFeedflowItemAdgroupAdzonePageResultDto.Put(v)
}
