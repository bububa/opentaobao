package ticket

// Scenic 
/* model for simplify = false
type Scenic struct {

    // 标准景点ID
    
    AliScenicId   string `json:"ali_scenic_id,omitempty"`
    

    // 标准景点名称
    
    AliScenicName   string `json:"ali_scenic_name,omitempty"`
    

    // 商家景点ID
    
    OutScenicId   string `json:"out_scenic_id,omitempty"`
    

    // 商家景点名称
    
    OutScenicName   string `json:"out_scenic_name,omitempty"`
    

    // 收费项目列表
    
    ProductList  struct {
        Product  []Product `json:"product,omitempty"`
    } `json:"product_list,omitempty"`
    

}
*/

// Scenic 
type Scenic struct {

    // 标准景点ID
    AliScenicId   string `json:"ali_scenic_id,omitempty"`

    // 标准景点名称
    AliScenicName   string `json:"ali_scenic_name,omitempty"`

    // 商家景点ID
    OutScenicId   string `json:"out_scenic_id,omitempty"`

    // 商家景点名称
    OutScenicName   string `json:"out_scenic_name,omitempty"`

    // 收费项目列表
    ProductList   []Product `json:"product_list,omitempty"`

}
