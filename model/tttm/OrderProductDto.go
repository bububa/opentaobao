package tttm

import (
	"sync"
)

// OrderProductDto 结构体
type OrderProductDto struct {
	// 货品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 货品价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 货品数量
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 货品编码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
}

var poolOrderProductDto = sync.Pool{
	New: func() any {
		return new(OrderProductDto)
	},
}

// GetOrderProductDto() 从对象池中获取OrderProductDto
func GetOrderProductDto() *OrderProductDto {
	return poolOrderProductDto.Get().(*OrderProductDto)
}

// ReleaseOrderProductDto 释放OrderProductDto
func ReleaseOrderProductDto(v *OrderProductDto) {
	v.ProductName = ""
	v.Price = ""
	v.Amount = ""
	v.ProductCode = ""
	poolOrderProductDto.Put(v)
}
