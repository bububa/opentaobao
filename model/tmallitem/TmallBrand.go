package tmallitem

// TmallBrand 
type TmallBrand struct {
    // 搜索品牌id
    BrandId   int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 搜索品牌名字
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
}
