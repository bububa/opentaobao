package koubeimall

// ItemGroupContent 
type ItemGroupContent struct {

    // 组名
    
    GroupName   string `json:"group_name,omitempty" xml:"group_name,omitempty"`
    

    // 详情组列表
    
    ContentGroupDetailList   []ItemGroupContentDetail `json:"content_group_detail_list,omitempty" xml:"content_group_detail_list,omitempty"`
    

}
