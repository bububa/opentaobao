package lsticitem

import (
	"sync"
)

// Stock 结构体
type Stock struct {
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 仓库名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 可售库存数量
	StockAmount int64 `json:"stock_amount,omitempty" xml:"stock_amount,omitempty"`
}

var poolStock = sync.Pool{
	New: func() any {
		return new(Stock)
	},
}

// GetStock() 从对象池中获取Stock
func GetStock() *Stock {
	return poolStock.Get().(*Stock)
}

// ReleaseStock 释放Stock
func ReleaseStock(v *Stock) {
	v.WarehouseCode = ""
	v.WarehouseName = ""
	v.StockAmount = 0
	poolStock.Put(v)
}
