package lstwarehouse

import (
	"sync"
)

// Stocklist 结构体
type Stocklist struct {
	// 商家仓仓库code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 要设置的库存数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolStocklist = sync.Pool{
	New: func() any {
		return new(Stocklist)
	},
}

// GetStocklist() 从对象池中获取Stocklist
func GetStocklist() *Stocklist {
	return poolStocklist.Get().(*Stocklist)
}

// ReleaseStocklist 释放Stocklist
func ReleaseStocklist(v *Stocklist) {
	v.WarehouseCode = ""
	v.Amount = 0
	v.ItemId = 0
	poolStocklist.Put(v)
}
