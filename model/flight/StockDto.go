package flight

import (
	"sync"
)

// StockDto 结构体
type StockDto struct {
	// 退票是否还库存
	ReturnStock int64 `json:"return_stock,omitempty" xml:"return_stock,omitempty"`
	// 库存限制
	StockLimit int64 `json:"stock_limit,omitempty" xml:"stock_limit,omitempty"`
	// 库存数量
	StockNum int64 `json:"stock_num,omitempty" xml:"stock_num,omitempty"`
}

var poolStockDto = sync.Pool{
	New: func() any {
		return new(StockDto)
	},
}

// GetStockDto() 从对象池中获取StockDto
func GetStockDto() *StockDto {
	return poolStockDto.Get().(*StockDto)
}

// ReleaseStockDto 释放StockDto
func ReleaseStockDto(v *StockDto) {
	v.ReturnStock = 0
	v.StockLimit = 0
	v.StockNum = 0
	poolStockDto.Put(v)
}
