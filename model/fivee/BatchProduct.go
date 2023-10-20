package fivee

import (
	"sync"
)

// BatchProduct 结构体
type BatchProduct struct {
	// 到期日期
	DueDate string `json:"due_date,omitempty" xml:"due_date,omitempty"`
	// 原产国
	OriginCountry string `json:"origin_country,omitempty" xml:"origin_country,omitempty"`
	// 生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
}

var poolBatchProduct = sync.Pool{
	New: func() any {
		return new(BatchProduct)
	},
}

// GetBatchProduct() 从对象池中获取BatchProduct
func GetBatchProduct() *BatchProduct {
	return poolBatchProduct.Get().(*BatchProduct)
}

// ReleaseBatchProduct 释放BatchProduct
func ReleaseBatchProduct(v *BatchProduct) {
	v.DueDate = ""
	v.OriginCountry = ""
	v.ProduceDate = ""
	poolBatchProduct.Put(v)
}
