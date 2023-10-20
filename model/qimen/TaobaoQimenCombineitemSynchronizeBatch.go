package qimen

import (
	"sync"
)

// TaobaoQimenCombineitemSynchronizeBatch 结构体
type TaobaoQimenCombineitemSynchronizeBatch struct {
	// test
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// test
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// test
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// test
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// test
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// test
	ActualQty string `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
	// test
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolTaobaoQimenCombineitemSynchronizeBatch = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemSynchronizeBatch)
	},
}

// GetTaobaoQimenCombineitemSynchronizeBatch() 从对象池中获取TaobaoQimenCombineitemSynchronizeBatch
func GetTaobaoQimenCombineitemSynchronizeBatch() *TaobaoQimenCombineitemSynchronizeBatch {
	return poolTaobaoQimenCombineitemSynchronizeBatch.Get().(*TaobaoQimenCombineitemSynchronizeBatch)
}

// ReleaseTaobaoQimenCombineitemSynchronizeBatch 释放TaobaoQimenCombineitemSynchronizeBatch
func ReleaseTaobaoQimenCombineitemSynchronizeBatch(v *TaobaoQimenCombineitemSynchronizeBatch) {
	v.BatchCode = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.ProduceCode = ""
	v.InventoryType = ""
	v.ActualQty = ""
	v.Quantity = ""
	poolTaobaoQimenCombineitemSynchronizeBatch.Put(v)
}
