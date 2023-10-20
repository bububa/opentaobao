package tblogistics

import (
	"sync"
)

// ItemTopDto 结构体
type ItemTopDto struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品类目
	ItemCate string `json:"item_cate,omitempty" xml:"item_cate,omitempty"`
	// 商品单价（原价）
	ItemValue int64 `json:"item_value,omitempty" xml:"item_value,omitempty"`
	// 商品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolItemTopDto = sync.Pool{
	New: func() any {
		return new(ItemTopDto)
	},
}

// GetItemTopDto() 从对象池中获取ItemTopDto
func GetItemTopDto() *ItemTopDto {
	return poolItemTopDto.Get().(*ItemTopDto)
}

// ReleaseItemTopDto 释放ItemTopDto
func ReleaseItemTopDto(v *ItemTopDto) {
	v.ItemName = ""
	v.ItemCate = ""
	v.ItemValue = 0
	v.ItemQuantity = 0
	v.ItemId = 0
	poolItemTopDto.Put(v)
}
