package hotel

// BrandType 
/* model for simplify = false
type BrandType struct {

    // 品牌列表
    
    BrandList  struct {
        Brand  []Brand `json:"brand,omitempty"`
    } `json:"brand_list,omitempty"`
    

    // 品牌类型
    
    TypeId   string `json:"type_id,omitempty"`
    

    // 品牌类型名称
    
    TypeName   string `json:"type_name,omitempty"`
    

}
*/

// BrandType 
type BrandType struct {

    // 品牌列表
    BrandList   []Brand `json:"brand_list,omitempty"`

    // 品牌类型
    TypeId   string `json:"type_id,omitempty"`

    // 品牌类型名称
    TypeName   string `json:"type_name,omitempty"`

}
