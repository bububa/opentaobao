package einvoice

import (
	"sync"
)

// BillItemDo 结构体
type BillItemDo struct {
	// 价税合计
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 规格型号,可选
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 商品单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 1 折扣行 2被折扣行 0普通行
	RowType int64 `json:"row_type,omitempty" xml:"row_type,omitempty"`
}

var poolBillItemDo = sync.Pool{
	New: func() any {
		return new(BillItemDo)
	},
}

// GetBillItemDo() 从对象池中获取BillItemDo
func GetBillItemDo() *BillItemDo {
	return poolBillItemDo.Get().(*BillItemDo)
}

// ReleaseBillItemDo 释放BillItemDo
func ReleaseBillItemDo(v *BillItemDo) {
	v.Amount = ""
	v.ItemName = ""
	v.Quantity = ""
	v.Specification = ""
	v.Unit = ""
	v.RowType = 0
	poolBillItemDo.Put(v)
}
