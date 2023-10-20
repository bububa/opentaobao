package drug

import (
	"sync"
)

// SuitSubItemDto 结构体
type SuitSubItemDto struct {
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 药标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商家自定义商品ID
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 个数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolSuitSubItemDto = sync.Pool{
	New: func() any {
		return new(SuitSubItemDto)
	},
}

// GetSuitSubItemDto() 从对象池中获取SuitSubItemDto
func GetSuitSubItemDto() *SuitSubItemDto {
	return poolSuitSubItemDto.Get().(*SuitSubItemDto)
}

// ReleaseSuitSubItemDto 释放SuitSubItemDto
func ReleaseSuitSubItemDto(v *SuitSubItemDto) {
	v.Unit = ""
	v.Title = ""
	v.OutId = ""
	v.ItemId = 0
	v.Quantity = 0
	poolSuitSubItemDto.Put(v)
}
