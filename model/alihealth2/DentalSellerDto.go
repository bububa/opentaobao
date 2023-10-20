package alihealth2

import (
	"sync"
)

// DentalSellerDto 结构体
type DentalSellerDto struct {
	// storeList
	StoreList []DentalStoreDto `json:"store_list,omitempty" xml:"store_list>dental_store_dto,omitempty"`
	// sellerName
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
}

var poolDentalSellerDto = sync.Pool{
	New: func() any {
		return new(DentalSellerDto)
	},
}

// GetDentalSellerDto() 从对象池中获取DentalSellerDto
func GetDentalSellerDto() *DentalSellerDto {
	return poolDentalSellerDto.Get().(*DentalSellerDto)
}

// ReleaseDentalSellerDto 释放DentalSellerDto
func ReleaseDentalSellerDto(v *DentalSellerDto) {
	v.StoreList = v.StoreList[:0]
	v.SellerName = ""
	poolDentalSellerDto.Put(v)
}
