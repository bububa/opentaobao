package campus

// SpaceTypeQuery 
type SpaceTypeQuery struct {

    // 每页大小
    
    Limit   int64 `json:"limit,omitempty" xml:"limit,omitempty"`
    

    // 大类id
    
    TopLevelId   int64 `json:"top_level_id,omitempty" xml:"top_level_id,omitempty"`
    

    // 中类id
    
    Pid   int64 `json:"pid,omitempty" xml:"pid,omitempty"`
    

    // 分组大类id
    
    GroupTopLevelId   int64 `json:"group_top_level_id,omitempty" xml:"group_top_level_id,omitempty"`
    

    // 小类id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 类别，室内，室外，逻辑
    
    Category   int64 `json:"category,omitempty" xml:"category,omitempty"`
    

    // 起始行
    
    StartRow   int64 `json:"start_row,omitempty" xml:"start_row,omitempty"`
    

    // 空间大类id
    
    SpaceTopLevelId   int64 `json:"space_top_level_id,omitempty" xml:"space_top_level_id,omitempty"`
    

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // 模糊查询key
    
    Key   string `json:"key,omitempty" xml:"key,omitempty"`
    

}
