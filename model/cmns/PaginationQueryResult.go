package cmns

// PaginationQueryResult 
/* model for simplify = false
type PaginationQueryResult struct {

    // ack记录列表
    
    List  struct {
        MessageAckResult  []MessageAckResult `json:"message_ack_result,omitempty"`
    } `json:"list,omitempty"`
    

    // 分页数据
    
    Pagination  *struct {
        Pagination  *Pagination `json:"pagination,omitempty"`
    } `json:"pagination,omitempty"`
    

}
*/

// PaginationQueryResult 
type PaginationQueryResult struct {

    // ack记录列表
    List   []MessageAckResult `json:"list,omitempty"`

    // 分页数据
    Pagination   *Pagination `json:"pagination,omitempty"`

}
