package tttm

import (
	"sync"
)

// ProductInfoDto 结构体
type ProductInfoDto struct {
	// 货品编码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 货品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 货品状态
	ProductStatus int64 `json:"product_status,omitempty" xml:"product_status,omitempty"`
	// 套餐数量
	SetAmount int64 `json:"set_amount,omitempty" xml:"set_amount,omitempty"`
	// 总库存
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 增量库存
	IncrementAmount int64 `json:"increment_amount,omitempty" xml:"increment_amount,omitempty"`
	// 出入库
	IncrementType int64 `json:"increment_type,omitempty" xml:"increment_type,omitempty"`
}

var poolProductInfoDto = sync.Pool{
	New: func() any {
		return new(ProductInfoDto)
	},
}

// GetProductInfoDto() 从对象池中获取ProductInfoDto
func GetProductInfoDto() *ProductInfoDto {
	return poolProductInfoDto.Get().(*ProductInfoDto)
}

// ReleaseProductInfoDto 释放ProductInfoDto
func ReleaseProductInfoDto(v *ProductInfoDto) {
	v.ProductCode = ""
	v.ProductName = ""
	v.ProductStatus = 0
	v.SetAmount = 0
	v.TotalAmount = 0
	v.IncrementAmount = 0
	v.IncrementType = 0
	poolProductInfoDto.Put(v)
}
