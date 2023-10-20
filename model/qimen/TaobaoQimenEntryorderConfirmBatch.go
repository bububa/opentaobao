package qimen

import (
	"sync"
)

// TaobaoQimenEntryorderConfirmBatch 结构体
type TaobaoQimenEntryorderConfirmBatch struct {
	// 批次编号
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 生产日期(YYYY-MM-DD)
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 过期日期(YYYY-MM-DD)
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 生产批号
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// 库存类型(ZP=正品;CC=残次;JS=机损;XS= 箱损;默认为ZP;)
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 实收数量(要求batchs节点下所有的实收数量之和等于orderline中的实收数量)
	ActualQty int64 `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
}

var poolTaobaoQimenEntryorderConfirmBatch = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderConfirmBatch)
	},
}

// GetTaobaoQimenEntryorderConfirmBatch() 从对象池中获取TaobaoQimenEntryorderConfirmBatch
func GetTaobaoQimenEntryorderConfirmBatch() *TaobaoQimenEntryorderConfirmBatch {
	return poolTaobaoQimenEntryorderConfirmBatch.Get().(*TaobaoQimenEntryorderConfirmBatch)
}

// ReleaseTaobaoQimenEntryorderConfirmBatch 释放TaobaoQimenEntryorderConfirmBatch
func ReleaseTaobaoQimenEntryorderConfirmBatch(v *TaobaoQimenEntryorderConfirmBatch) {
	v.BatchCode = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.ProduceCode = ""
	v.InventoryType = ""
	v.ActualQty = 0
	poolTaobaoQimenEntryorderConfirmBatch.Put(v)
}
