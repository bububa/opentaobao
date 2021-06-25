package cmns

// PaginationQueryResult 
type PaginationQueryResult struct {

    // ack记录列表
    List   []MessageAckResult `json:"list,omitempty"`

    // 分页数据
    Pagination   *Pagination `json:"pagination,omitempty"`

}
