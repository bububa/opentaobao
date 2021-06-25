package promotion

// PageInfo 
type PageInfo struct {

    // 第几页
    PageNum   int64 `json:"page_num,omitempty"`

    // 每页条数
    PageSize   int64 `json:"page_size,omitempty"`

    // 总共多少页
    Pages   int64 `json:"pages,omitempty"`

    // 总共多少条
    Total   int64 `json:"total,omitempty"`

}
