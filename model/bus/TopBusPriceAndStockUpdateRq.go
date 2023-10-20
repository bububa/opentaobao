package bus

import (
	"sync"
)

// TopBusPriceAndStockUpdateRq 结构体
type TopBusPriceAndStockUpdateRq struct {
	// 车次列表
	Numbers []BusNumberInfoDto `json:"numbers,omitempty" xml:"numbers>bus_number_info_dto,omitempty"`
}

var poolTopBusPriceAndStockUpdateRq = sync.Pool{
	New: func() any {
		return new(TopBusPriceAndStockUpdateRq)
	},
}

// GetTopBusPriceAndStockUpdateRq() 从对象池中获取TopBusPriceAndStockUpdateRq
func GetTopBusPriceAndStockUpdateRq() *TopBusPriceAndStockUpdateRq {
	return poolTopBusPriceAndStockUpdateRq.Get().(*TopBusPriceAndStockUpdateRq)
}

// ReleaseTopBusPriceAndStockUpdateRq 释放TopBusPriceAndStockUpdateRq
func ReleaseTopBusPriceAndStockUpdateRq(v *TopBusPriceAndStockUpdateRq) {
	v.Numbers = v.Numbers[:0]
	poolTopBusPriceAndStockUpdateRq.Put(v)
}
