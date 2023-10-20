package perfect

import (
	"sync"
)

// PerfectSalePropertyDto 结构体
type PerfectSalePropertyDto struct {
	// 属性key
	SalePropertyKey string `json:"sale_property_key,omitempty" xml:"sale_property_key,omitempty"`
	// 自定义属性值
	SalePropertyValue string `json:"sale_property_value,omitempty" xml:"sale_property_value,omitempty"`
}

var poolPerfectSalePropertyDto = sync.Pool{
	New: func() any {
		return new(PerfectSalePropertyDto)
	},
}

// GetPerfectSalePropertyDto() 从对象池中获取PerfectSalePropertyDto
func GetPerfectSalePropertyDto() *PerfectSalePropertyDto {
	return poolPerfectSalePropertyDto.Get().(*PerfectSalePropertyDto)
}

// ReleasePerfectSalePropertyDto 释放PerfectSalePropertyDto
func ReleasePerfectSalePropertyDto(v *PerfectSalePropertyDto) {
	v.SalePropertyKey = ""
	v.SalePropertyValue = ""
	poolPerfectSalePropertyDto.Put(v)
}
