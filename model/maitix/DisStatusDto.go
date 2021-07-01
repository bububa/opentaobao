package maitix

// DisStatusDto 结构体
type DisStatusDto struct {
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 1-可售 0不可售
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 票品状态列表
	DisTicketItemStatusDTOList []DisTicketItemStatusDto `json:"dis_ticket_item_status_d_t_o_list,omitempty" xml:"dis_ticket_item_status_d_t_o_list>dis_ticket_item_status_dto,omitempty"`
	// 场次状态列表
	DisPerformStatusDTOList []DisPerformStatusDto `json:"dis_perform_status_d_t_o_list,omitempty" xml:"dis_perform_status_d_t_o_list>dis_perform_status_dto,omitempty"`
}
