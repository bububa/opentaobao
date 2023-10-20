package ticket

import (
	"sync"
)

// DateInventory 结构体
type DateInventory struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 日期级别自定义商家编码，为该sku下每一天都设置一个自定义商家编码。如果outSkuDateId为空，则该天的商家自定义编码将以outSkuId为准
	OutSkuDateId string `json:"out_sku_date_id,omitempty" xml:"out_sku_date_id,omitempty"`
	// 价格，精确到分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
}

var poolDateInventory = sync.Pool{
	New: func() any {
		return new(DateInventory)
	},
}

// GetDateInventory() 从对象池中获取DateInventory
func GetDateInventory() *DateInventory {
	return poolDateInventory.Get().(*DateInventory)
}

// ReleaseDateInventory 释放DateInventory
func ReleaseDateInventory(v *DateInventory) {
	v.Date = ""
	v.OutSkuDateId = ""
	v.Price = 0
	v.Stock = 0
	poolDateInventory.Put(v)
}
