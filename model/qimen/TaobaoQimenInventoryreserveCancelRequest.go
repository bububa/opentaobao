package qimen

import (
	"sync"
)

// TaobaoQimenInventoryreserveCancelRequest 结构体
type TaobaoQimenInventoryreserveCancelRequest struct {
	// 奇门仓储字段
	ItemInventories []ItemInventory `json:"itemInventories,omitempty" xml:"itemInventories>item_inventory,omitempty"`
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 奇门仓储字段
	OrderSource string `json:"orderSource,omitempty" xml:"orderSource,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenInventoryreserveCancelMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolTaobaoQimenInventoryreserveCancelRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryreserveCancelRequest)
	},
}

// GetTaobaoQimenInventoryreserveCancelRequest() 从对象池中获取TaobaoQimenInventoryreserveCancelRequest
func GetTaobaoQimenInventoryreserveCancelRequest() *TaobaoQimenInventoryreserveCancelRequest {
	return poolTaobaoQimenInventoryreserveCancelRequest.Get().(*TaobaoQimenInventoryreserveCancelRequest)
}

// ReleaseTaobaoQimenInventoryreserveCancelRequest 释放TaobaoQimenInventoryreserveCancelRequest
func ReleaseTaobaoQimenInventoryreserveCancelRequest(v *TaobaoQimenInventoryreserveCancelRequest) {
	v.ItemInventories = v.ItemInventories[:0]
	v.OwnerCode = ""
	v.OrderCode = ""
	v.OrderSource = ""
	v.ExtendProps = nil
	poolTaobaoQimenInventoryreserveCancelRequest.Put(v)
}
