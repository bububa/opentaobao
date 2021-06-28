package tmc

// Paginator 
type Paginator struct {

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 当前第几页
    
    CurrentPageNum   int64 `json:"current_page_num,omitempty" xml:"current_page_num,omitempty"`
    

}
