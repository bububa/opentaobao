package campus

// Page 
type Page struct {

    // 总计数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 返回内容
    
    Items   []UserRolesDto `json:"items,omitempty" xml:"items,omitempty"`
    

    // 每页大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    

    // 每页多少条
    
    Limit   int64 `json:"limit,omitempty" xml:"limit,omitempty"`
    

    // 总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

    // 分页结果集合
    
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
    

    // 分页集合
    
    ResultList   []DeviceStandardApiDto `json:"result_list,omitempty" xml:"result_list,omitempty"`
    

}
