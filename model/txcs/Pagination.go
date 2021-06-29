package txcs

// Pagination 
type Pagination struct {

    // 总数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 单页个数
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

}
