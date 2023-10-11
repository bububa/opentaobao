package simba

// ShopCategoryVo 结构体
type ShopCategoryVo struct {
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 父级类目id
	ParentCategoryId int64 `json:"parent_category_id,omitempty" xml:"parent_category_id,omitempty"`
}
