package logistic

import (
	"sync"
)

// OrderWarehouseRouteGetItems 结构体
type OrderWarehouseRouteGetItems struct {
	// 商品信息
	Item *OrderWarehouseRouteGetItem `json:"item,omitempty" xml:"item,omitempty"`
}

var poolOrderWarehouseRouteGetItems = sync.Pool{
	New: func() any {
		return new(OrderWarehouseRouteGetItems)
	},
}

// GetOrderWarehouseRouteGetItems() 从对象池中获取OrderWarehouseRouteGetItems
func GetOrderWarehouseRouteGetItems() *OrderWarehouseRouteGetItems {
	return poolOrderWarehouseRouteGetItems.Get().(*OrderWarehouseRouteGetItems)
}

// ReleaseOrderWarehouseRouteGetItems 释放OrderWarehouseRouteGetItems
func ReleaseOrderWarehouseRouteGetItems(v *OrderWarehouseRouteGetItems) {
	v.Item = nil
	poolOrderWarehouseRouteGetItems.Put(v)
}
