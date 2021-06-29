package tuanhotel

// TopRootShopCategoryVoList 
type TopRootShopCategoryVoList struct {
    // 二级类目列表
    ShopCategoryList   []ShopCategoryVoList `json:"shop_category_list,omitempty" xml:"shop_category_list>shop_category_vo_list,omitempty"`
    // 一级类目名称
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    // 一级类目ID
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
