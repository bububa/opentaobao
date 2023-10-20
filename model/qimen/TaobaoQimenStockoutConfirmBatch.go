package qimen

import (
	"sync"
)

// TaobaoQimenStockoutConfirmBatch 结构体
type TaobaoQimenStockoutConfirmBatch struct {
	// 批次编号
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 生产日期
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 过期日期
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 生产批号
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// 库存类型
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 实发数量
	ActualQty int64 `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
}

var poolTaobaoQimenStockoutConfirmBatch = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockoutConfirmBatch)
	},
}

// GetTaobaoQimenStockoutConfirmBatch() 从对象池中获取TaobaoQimenStockoutConfirmBatch
func GetTaobaoQimenStockoutConfirmBatch() *TaobaoQimenStockoutConfirmBatch {
	return poolTaobaoQimenStockoutConfirmBatch.Get().(*TaobaoQimenStockoutConfirmBatch)
}

// ReleaseTaobaoQimenStockoutConfirmBatch 释放TaobaoQimenStockoutConfirmBatch
func ReleaseTaobaoQimenStockoutConfirmBatch(v *TaobaoQimenStockoutConfirmBatch) {
	v.BatchCode = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.ProduceCode = ""
	v.InventoryType = ""
	v.ActualQty = 0
	poolTaobaoQimenStockoutConfirmBatch.Put(v)
}
