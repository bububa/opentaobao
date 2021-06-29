package aecreatives

// Category 
type Category struct {
    // 类目ID
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    // 类目名称
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    // 父类目ID
    ParentCategoryId   int64 `json:"parent_category_id,omitempty" xml:"parent_category_id,omitempty"`
}
