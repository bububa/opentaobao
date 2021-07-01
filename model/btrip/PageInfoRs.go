package btrip

// PageInfoRS 
type PageInfoRS struct {
    // 当前页
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    // 每页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 总记录数
    TotalNumber   int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}
