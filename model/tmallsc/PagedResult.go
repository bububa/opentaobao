package tmallsc

// PagedResult 
type PagedResult struct {
    // 总条数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 每页条数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 结算明细list
    DataList   []Datalist `json:"data_list,omitempty" xml:"data_list>datalist,omitempty"`
}
