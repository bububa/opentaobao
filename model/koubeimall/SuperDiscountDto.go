package koubeimall

import (
	"sync"
)

// SuperDiscountDto 结构体
type SuperDiscountDto struct {
	// 商品信息list
	ItemList []ItemDto `json:"item_list,omitempty" xml:"item_list>item_dto,omitempty"`
}

var poolSuperDiscountDto = sync.Pool{
	New: func() any {
		return new(SuperDiscountDto)
	},
}

// GetSuperDiscountDto() 从对象池中获取SuperDiscountDto
func GetSuperDiscountDto() *SuperDiscountDto {
	return poolSuperDiscountDto.Get().(*SuperDiscountDto)
}

// ReleaseSuperDiscountDto 释放SuperDiscountDto
func ReleaseSuperDiscountDto(v *SuperDiscountDto) {
	v.ItemList = v.ItemList[:0]
	poolSuperDiscountDto.Put(v)
}
