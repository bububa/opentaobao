package moscm

// Paginator 
type Paginator struct {

    // 当前页码
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // 每页显示多少条记录
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

}
