package tuanhotel

// ShopCategoryVoList 
type ShopCategoryVoList struct {
    // 二级类目名称
    CategoriesName   string `json:"categories_name,omitempty" xml:"categories_name,omitempty"`
    // 二级类目ID
    CategoriesId   int64 `json:"categories_id,omitempty" xml:"categories_id,omitempty"`
}
