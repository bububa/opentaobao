package fenxiao

// QueryPagination 
type QueryPagination struct {
    // 当前页码数
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    // 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
