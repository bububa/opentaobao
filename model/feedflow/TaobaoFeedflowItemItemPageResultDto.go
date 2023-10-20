package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemItemPageResultDto 结构体
type TaobaoFeedflowItemItemPageResultDto struct {
	// 商品列表
	ItemList []ItemDto `json:"item_list,omitempty" xml:"item_list>item_dto,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 商品总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemItemPageResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemItemPageResultDto)
	},
}

// GetTaobaoFeedflowItemItemPageResultDto() 从对象池中获取TaobaoFeedflowItemItemPageResultDto
func GetTaobaoFeedflowItemItemPageResultDto() *TaobaoFeedflowItemItemPageResultDto {
	return poolTaobaoFeedflowItemItemPageResultDto.Get().(*TaobaoFeedflowItemItemPageResultDto)
}

// ReleaseTaobaoFeedflowItemItemPageResultDto 释放TaobaoFeedflowItemItemPageResultDto
func ReleaseTaobaoFeedflowItemItemPageResultDto(v *TaobaoFeedflowItemItemPageResultDto) {
	v.ItemList = v.ItemList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.Success = false
	poolTaobaoFeedflowItemItemPageResultDto.Put(v)
}
