package alihealth2

import (
	"sync"
)

// DentalStoreDto 结构体
type DentalStoreDto struct {
	// itemList
	ItemList []DentalItemDto `json:"item_list,omitempty" xml:"item_list>dental_item_dto,omitempty"`
	// storeName
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// storeId
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolDentalStoreDto = sync.Pool{
	New: func() any {
		return new(DentalStoreDto)
	},
}

// GetDentalStoreDto() 从对象池中获取DentalStoreDto
func GetDentalStoreDto() *DentalStoreDto {
	return poolDentalStoreDto.Get().(*DentalStoreDto)
}

// ReleaseDentalStoreDto 释放DentalStoreDto
func ReleaseDentalStoreDto(v *DentalStoreDto) {
	v.ItemList = v.ItemList[:0]
	v.StoreName = ""
	v.StoreId = 0
	poolDentalStoreDto.Put(v)
}
