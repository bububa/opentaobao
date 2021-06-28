package category

// CategoryDto 
/* model for simplify = false
type CategoryDto struct {

    // 子节点
    
    Childrens  struct {
        CategoryDto  []CategoryDto `json:"category_dto,omitempty"`
    } `json:"childrens,omitempty"`
    

    // 类目名称
    
    CategoryName   string `json:"category_name,omitempty"`
    

    // 类目ID
    
    CategoryId   string `json:"category_id,omitempty"`
    

}
*/

// CategoryDto 
type CategoryDto struct {

    // 子节点
    Childrens   []CategoryDto `json:"childrens,omitempty"`

    // 类目名称
    CategoryName   string `json:"category_name,omitempty"`

    // 类目ID
    CategoryId   string `json:"category_id,omitempty"`

}
