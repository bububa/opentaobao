package wdk

// Pagination 
type Pagination struct {

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 总记录数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 页容量
    PageSize   int64 `json:"page_size,omitempty"`

    // 当前页码
    CurrentPage   int64 `json:"current_page,omitempty"`

}
