package tblogistics

// TaobaologisticsinstanttracesearchResult 结构体
type TaobaologisticsinstanttracesearchResult struct {
	// 运单列表
	MailList []TopLogisticsMailDto `json:"mail_list,omitempty" xml:"mail_list>top_logistics_mail_dto,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
