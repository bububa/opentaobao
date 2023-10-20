package rhino

import (
	"sync"
)

// ClothingSkuDto 结构体
type ClothingSkuDto struct {
	// 成衣物料名称ItemName
	StyleName string `json:"style_name,omitempty" xml:"style_name,omitempty"`
	// 成衣sku编码Item
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 成衣rfid编码
	Rfid string `json:"rfid,omitempty" xml:"rfid,omitempty"`
	// 发货数量item_count
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolClothingSkuDto = sync.Pool{
	New: func() any {
		return new(ClothingSkuDto)
	},
}

// GetClothingSkuDto() 从对象池中获取ClothingSkuDto
func GetClothingSkuDto() *ClothingSkuDto {
	return poolClothingSkuDto.Get().(*ClothingSkuDto)
}

// ReleaseClothingSkuDto 释放ClothingSkuDto
func ReleaseClothingSkuDto(v *ClothingSkuDto) {
	v.StyleName = ""
	v.Id = ""
	v.Rfid = ""
	v.Amount = 0
	poolClothingSkuDto.Put(v)
}
