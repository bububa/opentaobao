package tttm

import (
	"sync"
)

// StockInfoDto 结构体
type StockInfoDto struct {
	// 工厂仓
	FactoryDepot []ProductInfoDto `json:"factory_depot,omitempty" xml:"factory_depot>product_info_dto,omitempty"`
	// 电商仓
	ShopDepot []ProductInfoDto `json:"shop_depot,omitempty" xml:"shop_depot>product_info_dto,omitempty"`
}

var poolStockInfoDto = sync.Pool{
	New: func() any {
		return new(StockInfoDto)
	},
}

// GetStockInfoDto() 从对象池中获取StockInfoDto
func GetStockInfoDto() *StockInfoDto {
	return poolStockInfoDto.Get().(*StockInfoDto)
}

// ReleaseStockInfoDto 释放StockInfoDto
func ReleaseStockInfoDto(v *StockInfoDto) {
	v.FactoryDepot = v.FactoryDepot[:0]
	v.ShopDepot = v.ShopDepot[:0]
	poolStockInfoDto.Put(v)
}
