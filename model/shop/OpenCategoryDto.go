package shop

// OpenCategoryDto 结构体
type OpenCategoryDto struct {
	// 子分类对象
	SubShopCategoryList []OpenCategoryDto `json:"sub_shop_category_list,omitempty" xml:"sub_shop_category_list>open_category_dto,omitempty"`
	// 分类名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 分类主图
	CategoryImage string `json:"category_image,omitempty" xml:"category_image,omitempty"`
	// 分类id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
