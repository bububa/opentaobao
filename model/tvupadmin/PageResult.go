package tvupadmin

// PageResult 
type PageResult struct {

    // 技能列表
    
    Results   []SkillSimpleView `json:"results,omitempty" xml:"results,omitempty"`
    

    // 总页数
    
    PageCount   int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
    

    // 分页单位
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

    // 当前页
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

}
