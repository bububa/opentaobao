package maitix

// DisPerformStatusQueryDto 
type DisPerformStatusQueryDto struct {
    // 场次id
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    // 是否查询场次对应的票品状态
    QueryTicketItemStatus   bool `json:"query_ticket_item_status,omitempty" xml:"query_ticket_item_status,omitempty"`
}
