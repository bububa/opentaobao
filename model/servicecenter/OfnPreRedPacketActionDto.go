package servicecenter

// OfnPreRedPacketActionDto 结构体
type OfnPreRedPacketActionDto struct {
	// 资金池的记录
	AfterFundRecordList []OfnPreRedPacketFundRecordDto `json:"after_fund_record_list,omitempty" xml:"after_fund_record_list>ofn_pre_red_packet_fund_record_dto,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 操作类型。1=天猫发预付红包；2=天猫发尾款红包；3=天猫扣回红包；4=回收商扣回红包
	ActionType int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 状态。初始化=1，重试中=2，失败=3，成功=4
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
