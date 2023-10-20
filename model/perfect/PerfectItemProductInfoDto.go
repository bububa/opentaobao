package perfect

import (
	"sync"
)

// PerfectItemProductInfoDto 结构体
type PerfectItemProductInfoDto struct {
	// 品牌ID
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 叶子类目ID
	CategoryCode string `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// 货号
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
}

var poolPerfectItemProductInfoDto = sync.Pool{
	New: func() any {
		return new(PerfectItemProductInfoDto)
	},
}

// GetPerfectItemProductInfoDto() 从对象池中获取PerfectItemProductInfoDto
func GetPerfectItemProductInfoDto() *PerfectItemProductInfoDto {
	return poolPerfectItemProductInfoDto.Get().(*PerfectItemProductInfoDto)
}

// ReleasePerfectItemProductInfoDto 释放PerfectItemProductInfoDto
func ReleasePerfectItemProductInfoDto(v *PerfectItemProductInfoDto) {
	v.BrandCode = ""
	v.CategoryCode = ""
	v.ProductCode = ""
	poolPerfectItemProductInfoDto.Put(v)
}
