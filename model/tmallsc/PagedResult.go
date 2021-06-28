package tmallsc

// PagedResult 
/* model for simplify = false
type PagedResult struct {

    // 总条数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 每页条数
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 结算明细list
    
    DataList  struct {
        Datalist  []Datalist `json:"datalist,omitempty"`
    } `json:"data_list,omitempty"`
    

}
*/

// PagedResult 
type PagedResult struct {

    // 总条数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 每页条数
    PageSize   int64 `json:"page_size,omitempty"`

    // 结算明细list
    DataList   []Datalist `json:"data_list,omitempty"`

}
