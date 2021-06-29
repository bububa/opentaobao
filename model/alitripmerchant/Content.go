package alitripmerchant

// Content 
type Content struct {

    // 高档型
    
    Upscales   []BrandListDto `json:"upscales,omitempty" xml:"upscales,omitempty"`
    

    // 豪华型
    
    Luxurys   []BrandListDto `json:"luxurys,omitempty" xml:"luxurys,omitempty"`
    

    // 舒适型
    
    CmForTables   []BrandListDto `json:"cm_for_tables,omitempty" xml:"cm_for_tables,omitempty"`
    

    // 经济型
    
    Economics   []BrandListDto `json:"economics,omitempty" xml:"economics,omitempty"`
    

}
