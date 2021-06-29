package moscm

// PagedList 
type PagedList struct {
    // 当前页
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    // 数据结果集
    List   []Cspudto `json:"list,omitempty" xml:"list>cspudto,omitempty"`
    // 每页获取多少条记录
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 总记录数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
