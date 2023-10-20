package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemAdzoneListResultDto 结构体
type TaobaoFeedflowItemAdzoneListResultDto struct {
	// 广告位列表
	AdzoneList []AdzoneDto `json:"adzone_list,omitempty" xml:"adzone_list>adzone_dto,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemAdzoneListResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdzoneListResultDto)
	},
}

// GetTaobaoFeedflowItemAdzoneListResultDto() 从对象池中获取TaobaoFeedflowItemAdzoneListResultDto
func GetTaobaoFeedflowItemAdzoneListResultDto() *TaobaoFeedflowItemAdzoneListResultDto {
	return poolTaobaoFeedflowItemAdzoneListResultDto.Get().(*TaobaoFeedflowItemAdzoneListResultDto)
}

// ReleaseTaobaoFeedflowItemAdzoneListResultDto 释放TaobaoFeedflowItemAdzoneListResultDto
func ReleaseTaobaoFeedflowItemAdzoneListResultDto(v *TaobaoFeedflowItemAdzoneListResultDto) {
	v.AdzoneList = v.AdzoneList[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemAdzoneListResultDto.Put(v)
}
