package product

import (
	"sync"
)

// AfterSale 结构体
type AfterSale struct {
	// 名称
	AfterSaleName string `json:"after_sale_name,omitempty" xml:"after_sale_name,omitempty"`
	// tfs地址
	AfterSalePath string `json:"after_sale_path,omitempty" xml:"after_sale_path,omitempty"`
	// id
	AfterSaleId int64 `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty"`
}

var poolAfterSale = sync.Pool{
	New: func() any {
		return new(AfterSale)
	},
}

// GetAfterSale() 从对象池中获取AfterSale
func GetAfterSale() *AfterSale {
	return poolAfterSale.Get().(*AfterSale)
}

// ReleaseAfterSale 释放AfterSale
func ReleaseAfterSale(v *AfterSale) {
	v.AfterSaleName = ""
	v.AfterSalePath = ""
	v.AfterSaleId = 0
	poolAfterSale.Put(v)
}
