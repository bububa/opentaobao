package qimen

import (
	"sync"
)

// ItemsSynchronizeRequest 结构体
type ItemsSynchronizeRequest struct {
	// 同步的商品信息
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 操作类型(两种类型：add|update)
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenItemsSynchronizeMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolItemsSynchronizeRequest = sync.Pool{
	New: func() any {
		return new(ItemsSynchronizeRequest)
	},
}

// GetItemsSynchronizeRequest() 从对象池中获取ItemsSynchronizeRequest
func GetItemsSynchronizeRequest() *ItemsSynchronizeRequest {
	return poolItemsSynchronizeRequest.Get().(*ItemsSynchronizeRequest)
}

// ReleaseItemsSynchronizeRequest 释放ItemsSynchronizeRequest
func ReleaseItemsSynchronizeRequest(v *ItemsSynchronizeRequest) {
	v.Items = v.Items[:0]
	v.ActionType = ""
	v.WarehouseCode = ""
	v.OwnerCode = ""
	v.ExtendProps = nil
	poolItemsSynchronizeRequest.Put(v)
}
