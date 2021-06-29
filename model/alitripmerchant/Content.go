package alitripmerchant

// Content 
type Content struct {
    // 高档型
    Upscales   []BrandListDTO `json:"upscales,omitempty" xml:"upscales>brand_list_dto,omitempty"`
    // 豪华型
    Luxurys   []BrandListDTO `json:"luxurys,omitempty" xml:"luxurys>brand_list_dto,omitempty"`
    // 舒适型
    CmForTables   []BrandListDTO `json:"cm_for_tables,omitempty" xml:"cm_for_tables>brand_list_dto,omitempty"`
    // 经济型
    Economics   []BrandListDTO `json:"economics,omitempty" xml:"economics>brand_list_dto,omitempty"`
}
