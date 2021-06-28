package scbp

// CountryTagView 
/* model for simplify = false
type CountryTagView struct {

    // 最近7天效果数据
    
    Effect  *struct {
        Effect7d  *Effect7d `json:"effect7d,omitempty"`
    } `json:"effect,omitempty"`
    

    // 溢价百分比
    
    Discount   int64 `json:"discount,omitempty"`
    

    // 标签中文名
    
    TagName   string `json:"tag_name,omitempty"`
    

    // 标签id
    
    TagId   string `json:"tag_id,omitempty"`
    

}
*/

// CountryTagView 
type CountryTagView struct {

    // 最近7天效果数据
    Effect   *Effect7d `json:"effect,omitempty"`

    // 溢价百分比
    Discount   int64 `json:"discount,omitempty"`

    // 标签中文名
    TagName   string `json:"tag_name,omitempty"`

    // 标签id
    TagId   string `json:"tag_id,omitempty"`

}
