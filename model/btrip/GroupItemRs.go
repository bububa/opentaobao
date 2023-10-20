package btrip

import (
	"sync"
)

// GroupItemRs 结构体
type GroupItemRs struct {
	// 成人、儿童、老人对应的报价信息
	ShoppingItems []PassengerFlightShoppingItemRs `json:"shopping_items,omitempty" xml:"shopping_items>passenger_flight_shopping_item_rs,omitempty"`
	// 行程信息
	RouteInfo *RouteInfoRs `json:"route_info,omitempty" xml:"route_info,omitempty"`
}

var poolGroupItemRs = sync.Pool{
	New: func() any {
		return new(GroupItemRs)
	},
}

// GetGroupItemRs() 从对象池中获取GroupItemRs
func GetGroupItemRs() *GroupItemRs {
	return poolGroupItemRs.Get().(*GroupItemRs)
}

// ReleaseGroupItemRs 释放GroupItemRs
func ReleaseGroupItemRs(v *GroupItemRs) {
	v.ShoppingItems = v.ShoppingItems[:0]
	v.RouteInfo = nil
	poolGroupItemRs.Put(v)
}
