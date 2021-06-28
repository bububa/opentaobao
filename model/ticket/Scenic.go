package ticket

// Scenic 
type Scenic struct {

    // 标准景点ID
    
    AliScenicId   string `json:"ali_scenic_id,omitempty" xml:"ali_scenic_id,omitempty"`
    

    // 标准景点名称
    
    AliScenicName   string `json:"ali_scenic_name,omitempty" xml:"ali_scenic_name,omitempty"`
    

    // 商家景点ID
    
    OutScenicId   string `json:"out_scenic_id,omitempty" xml:"out_scenic_id,omitempty"`
    

    // 商家景点名称
    
    OutScenicName   string `json:"out_scenic_name,omitempty" xml:"out_scenic_name,omitempty"`
    

    // 收费项目列表
    
    ProductList   []Product `json:"product_list,omitempty" xml:"product_list,omitempty"`
    

}
