package category

// CategoryDTO 
type CategoryDTO struct {
    // 子节点
    Childrens   []CategoryDTO `json:"childrens,omitempty" xml:"childrens>category_dto,omitempty"`
    // 类目名称
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    // 类目ID
    CategoryId   string `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
