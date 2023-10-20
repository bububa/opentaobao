package alitripmerchant

import (
	"sync"
)

// Content 结构体
type Content struct {
	// 高档型
	Upscales []BrandListDto `json:"upscales,omitempty" xml:"upscales>brand_list_dto,omitempty"`
	// 豪华型
	Luxurys []BrandListDto `json:"luxurys,omitempty" xml:"luxurys>brand_list_dto,omitempty"`
	// 舒适型
	CmForTables []BrandListDto `json:"cm_for_tables,omitempty" xml:"cm_for_tables>brand_list_dto,omitempty"`
	// 经济型
	Economics []BrandListDto `json:"economics,omitempty" xml:"economics>brand_list_dto,omitempty"`
}

var poolContent = sync.Pool{
	New: func() any {
		return new(Content)
	},
}

// GetContent() 从对象池中获取Content
func GetContent() *Content {
	return poolContent.Get().(*Content)
}

// ReleaseContent 释放Content
func ReleaseContent(v *Content) {
	v.Upscales = v.Upscales[:0]
	v.Luxurys = v.Luxurys[:0]
	v.CmForTables = v.CmForTables[:0]
	v.Economics = v.Economics[:0]
	poolContent.Put(v)
}
