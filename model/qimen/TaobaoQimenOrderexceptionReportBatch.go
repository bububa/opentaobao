package qimen

import (
	"sync"
)

// TaobaoQimenOrderexceptionReportBatch 结构体
type TaobaoQimenOrderexceptionReportBatch struct {
	// 奇门仓储字段
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 奇门仓储字段
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 奇门仓储字段
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 奇门仓储字段
	ProduceCode string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
	// 奇门仓储字段
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 奇门仓储字段
	ActualQty string `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
	// 奇门仓储字段
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolTaobaoQimenOrderexceptionReportBatch = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderexceptionReportBatch)
	},
}

// GetTaobaoQimenOrderexceptionReportBatch() 从对象池中获取TaobaoQimenOrderexceptionReportBatch
func GetTaobaoQimenOrderexceptionReportBatch() *TaobaoQimenOrderexceptionReportBatch {
	return poolTaobaoQimenOrderexceptionReportBatch.Get().(*TaobaoQimenOrderexceptionReportBatch)
}

// ReleaseTaobaoQimenOrderexceptionReportBatch 释放TaobaoQimenOrderexceptionReportBatch
func ReleaseTaobaoQimenOrderexceptionReportBatch(v *TaobaoQimenOrderexceptionReportBatch) {
	v.BatchCode = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.ProduceCode = ""
	v.InventoryType = ""
	v.ActualQty = ""
	v.Quantity = ""
	poolTaobaoQimenOrderexceptionReportBatch.Put(v)
}
