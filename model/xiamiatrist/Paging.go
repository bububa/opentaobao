package xiamiatrist

// Paging 
type Paging struct {

    // 总页数
    
    Pages   int64 `json:"pages,omitempty" xml:"pages,omitempty"`
    

    // 总数
    
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    

    // 每页展示数
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 当前页码
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

}
