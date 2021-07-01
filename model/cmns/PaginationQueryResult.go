package cmns

// PaginationQueryResult 结构体
type PaginationQueryResult struct {
	// ack记录列表
	List []MessageAckResult `json:"list,omitempty" xml:"list>message_ack_result,omitempty"`
	// 分页数据
	Pagination *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`
}
