package alihealth2

import (
	"sync"
)

// DentalItemDto 结构体
type DentalItemDto struct {
	// itemName
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolDentalItemDto = sync.Pool{
	New: func() any {
		return new(DentalItemDto)
	},
}

// GetDentalItemDto() 从对象池中获取DentalItemDto
func GetDentalItemDto() *DentalItemDto {
	return poolDentalItemDto.Get().(*DentalItemDto)
}

// ReleaseDentalItemDto 释放DentalItemDto
func ReleaseDentalItemDto(v *DentalItemDto) {
	v.ItemName = ""
	v.ItemId = 0
	poolDentalItemDto.Put(v)
}
