package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdgroupCreativePageResultDto 结构体
type TaobaoFeedflowItemAdgroupCreativePageResultDto struct {
	// 绑定创意的列表
	CreativeBindList []CreativeBindDto `json:"creative_bind_list,omitempty" xml:"creative_bind_list>creative_bind_dto,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数目
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdgroupCreativePageResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupCreativePageResultDto)
	},
}

// GetTaobaoFeedflowItemAdgroupCreativePageResultDto() 从对象池中获取TaobaoFeedflowItemAdgroupCreativePageResultDto
func GetTaobaoFeedflowItemAdgroupCreativePageResultDto() *TaobaoFeedflowItemAdgroupCreativePageResultDto {
	return poolTaobaoFeedflowItemAdgroupCreativePageResultDto.Get().(*TaobaoFeedflowItemAdgroupCreativePageResultDto)
}

// ReleaseTaobaoFeedflowItemAdgroupCreativePageResultDto 释放TaobaoFeedflowItemAdgroupCreativePageResultDto
func ReleaseTaobaoFeedflowItemAdgroupCreativePageResultDto(v *TaobaoFeedflowItemAdgroupCreativePageResultDto) {
	v.CreativeBindList = v.CreativeBindList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.Success = false
	poolTaobaoFeedflowItemAdgroupCreativePageResultDto.Put(v)
}
