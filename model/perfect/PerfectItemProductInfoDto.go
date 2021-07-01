package perfect

// PerfectItemProductInfoDto 
type PerfectItemProductInfoDto struct {
    // 品牌ID
    BrandCode   string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
    // 叶子类目ID
    CategoryCode   string `json:"category_code,omitempty" xml:"category_code,omitempty"`
    // 货号
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
}
