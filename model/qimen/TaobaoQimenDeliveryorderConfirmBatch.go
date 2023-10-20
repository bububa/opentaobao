package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderConfirmBatch 结构体
type TaobaoQimenDeliveryorderConfirmBatch struct {
	// 批次编号
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 生产日期(YYYY-MM-DD)
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 过期日期(YYYY-MM-DD)
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 生产批号
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 货品sn编码
	SnCode string `json:"snCode,omitempty" xml:"snCode,omitempty"`
	// 实发数量(要求batchs节点下所有的实发数量之和等于orderline中的实发数量)
	ActualQty int64 `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
}

var poolTaobaoQimenDeliveryorderConfirmBatch = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderConfirmBatch)
	},
}

// GetTaobaoQimenDeliveryorderConfirmBatch() 从对象池中获取TaobaoQimenDeliveryorderConfirmBatch
func GetTaobaoQimenDeliveryorderConfirmBatch() *TaobaoQimenDeliveryorderConfirmBatch {
	return poolTaobaoQimenDeliveryorderConfirmBatch.Get().(*TaobaoQimenDeliveryorderConfirmBatch)
}

// ReleaseTaobaoQimenDeliveryorderConfirmBatch 释放TaobaoQimenDeliveryorderConfirmBatch
func ReleaseTaobaoQimenDeliveryorderConfirmBatch(v *TaobaoQimenDeliveryorderConfirmBatch) {
	v.BatchCode = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.ProduceCode = ""
	v.InventoryType = ""
	v.SnCode = ""
	v.ActualQty = 0
	poolTaobaoQimenDeliveryorderConfirmBatch.Put(v)
}
