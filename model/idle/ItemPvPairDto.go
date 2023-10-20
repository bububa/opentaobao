package idle

import (
	"sync"
)

// ItemPvPairDto 结构体
type ItemPvPairDto struct {
	// sku属性名
	PropText string `json:"prop_text,omitempty" xml:"prop_text,omitempty"`
	// sku属性值名称
	ValueText string `json:"value_text,omitempty" xml:"value_text,omitempty"`
}

var poolItemPvPairDto = sync.Pool{
	New: func() any {
		return new(ItemPvPairDto)
	},
}

// GetItemPvPairDto() 从对象池中获取ItemPvPairDto
func GetItemPvPairDto() *ItemPvPairDto {
	return poolItemPvPairDto.Get().(*ItemPvPairDto)
}

// ReleaseItemPvPairDto 释放ItemPvPairDto
func ReleaseItemPvPairDto(v *ItemPvPairDto) {
	v.PropText = ""
	v.ValueText = ""
	poolItemPvPairDto.Put(v)
}
