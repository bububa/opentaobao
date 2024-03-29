package qimen

import (
	"sync"
)

// TaobaoQimenReturnorderConfirmBatch 结构体
type TaobaoQimenReturnorderConfirmBatch struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 批次编号
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 生产日期(YYYY-MM-DD)
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 过期日期(YYYY-MM-DD)
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 生产批号
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;默认为ZP)
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 实收数量(要求batchs节点下所有的实收数量之和等于orderline中的实收数量)
	ActualQty int64 `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
}

var poolTaobaoQimenReturnorderConfirmBatch = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnorderConfirmBatch)
	},
}

// GetTaobaoQimenReturnorderConfirmBatch() 从对象池中获取TaobaoQimenReturnorderConfirmBatch
func GetTaobaoQimenReturnorderConfirmBatch() *TaobaoQimenReturnorderConfirmBatch {
	return poolTaobaoQimenReturnorderConfirmBatch.Get().(*TaobaoQimenReturnorderConfirmBatch)
}

// ReleaseTaobaoQimenReturnorderConfirmBatch 释放TaobaoQimenReturnorderConfirmBatch
func ReleaseTaobaoQimenReturnorderConfirmBatch(v *TaobaoQimenReturnorderConfirmBatch) {
	v.Remark = ""
	v.BatchCode = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.ProduceCode = ""
	v.InventoryType = ""
	v.ActualQty = 0
	poolTaobaoQimenReturnorderConfirmBatch.Put(v)
}
