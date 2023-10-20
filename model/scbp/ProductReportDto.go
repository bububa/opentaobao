package scbp

import (
	"sync"
)

// ProductReportDto 结构体
type ProductReportDto struct {
	// 返回实体集合
	ProductEffectList []ProductEffectDto `json:"product_effect_list,omitempty" xml:"product_effect_list>product_effect_dto,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolProductReportDto = sync.Pool{
	New: func() any {
		return new(ProductReportDto)
	},
}

// GetProductReportDto() 从对象池中获取ProductReportDto
func GetProductReportDto() *ProductReportDto {
	return poolProductReportDto.Get().(*ProductReportDto)
}

// ReleaseProductReportDto 释放ProductReportDto
func ReleaseProductReportDto(v *ProductReportDto) {
	v.ProductEffectList = v.ProductEffectList[:0]
	v.TotalCount = 0
	poolProductReportDto.Put(v)
}
