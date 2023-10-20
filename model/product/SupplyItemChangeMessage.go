package product

import (
	"sync"
)

// SupplyItemChangeMessage 结构体
type SupplyItemChangeMessage struct {
	// 货号列表
	ProductCodes []string `json:"product_codes,omitempty" xml:"product_codes>string,omitempty"`
	// 供应商门店，最大20个
	SupplyStoreId string `json:"supply_store_id,omitempty" xml:"supply_store_id,omitempty"`
}

var poolSupplyItemChangeMessage = sync.Pool{
	New: func() any {
		return new(SupplyItemChangeMessage)
	},
}

// GetSupplyItemChangeMessage() 从对象池中获取SupplyItemChangeMessage
func GetSupplyItemChangeMessage() *SupplyItemChangeMessage {
	return poolSupplyItemChangeMessage.Get().(*SupplyItemChangeMessage)
}

// ReleaseSupplyItemChangeMessage 释放SupplyItemChangeMessage
func ReleaseSupplyItemChangeMessage(v *SupplyItemChangeMessage) {
	v.ProductCodes = v.ProductCodes[:0]
	v.SupplyStoreId = ""
	poolSupplyItemChangeMessage.Put(v)
}
