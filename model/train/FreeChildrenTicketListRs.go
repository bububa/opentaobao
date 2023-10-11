package train

// FreeChildrenTicketListRs 结构体
type FreeChildrenTicketListRs struct {
	// 待处理列表
	FreeChildrenTicketNeedDealList []string `json:"free_children_ticket_need_deal_list,omitempty" xml:"free_children_ticket_need_deal_list>string,omitempty"`
}
