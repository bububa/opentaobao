package alitripmerchant

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
