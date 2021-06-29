package alihealthoutflow

// Page 
type Page struct {

    // 页对象
    
    Pages   []DrugDto `json:"pages,omitempty" xml:"pages,omitempty"`
    

    // 总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

    // 分页大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 页码
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

}
