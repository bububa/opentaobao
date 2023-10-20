package tmallservice

import (
	"sync"
)

// StoreOfferPriceDto 结构体
type StoreOfferPriceDto struct {
	// 服务code
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务skuCode
	ServiceSku string `json:"service_sku,omitempty" xml:"service_sku,omitempty"`
	// price
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolStoreOfferPriceDto = sync.Pool{
	New: func() any {
		return new(StoreOfferPriceDto)
	},
}

// GetStoreOfferPriceDto() 从对象池中获取StoreOfferPriceDto
func GetStoreOfferPriceDto() *StoreOfferPriceDto {
	return poolStoreOfferPriceDto.Get().(*StoreOfferPriceDto)
}

// ReleaseStoreOfferPriceDto 释放StoreOfferPriceDto
func ReleaseStoreOfferPriceDto(v *StoreOfferPriceDto) {
	v.ServiceCode = ""
	v.ServiceSku = ""
	v.Price = 0
	v.StoreId = 0
	poolStoreOfferPriceDto.Put(v)
}
