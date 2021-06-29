package eleenterpriserestaurant

// Activite 
type Activite struct {

    // 见附录【活动信息detail_type值】
    
    DetailType   int64 `json:"detail_type,omitempty" xml:"detail_type,omitempty"`
    

    // 活动名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 餐厅描述
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 活动ID
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

}
