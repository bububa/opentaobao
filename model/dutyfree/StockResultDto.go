package dutyfree

import (
	"sync"
)

// StockResultDto 结构体
type StockResultDto struct {
	// 条形码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
}

var poolStockResultDto = sync.Pool{
	New: func() any {
		return new(StockResultDto)
	},
}

// GetStockResultDto() 从对象池中获取StockResultDto
func GetStockResultDto() *StockResultDto {
	return poolStockResultDto.Get().(*StockResultDto)
}

// ReleaseStockResultDto 释放StockResultDto
func ReleaseStockResultDto(v *StockResultDto) {
	v.BarCode = ""
	v.Stock = 0
	poolStockResultDto.Put(v)
}
