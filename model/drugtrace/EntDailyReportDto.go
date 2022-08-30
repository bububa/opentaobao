package drugtrace

// EntDailyReportDto 结构体
type EntDailyReportDto struct {
	// 报告所有者名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 该报告的所有者唯一标识
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 报告开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 报告结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 其中生成的单据总数
	BillFileCount int64 `json:"bill_file_count,omitempty" xml:"bill_file_count,omitempty"`
	// 其中关联关系工生成包的数量
	RelationTaskCount int64 `json:"relation_task_count,omitempty" xml:"relation_task_count,omitempty"`
	// 其中生成的关联关系文件总数
	RelationFileCount int64 `json:"relation_file_count,omitempty" xml:"relation_file_count,omitempty"`
	// 其中已经接收到正确回执的单据包数量是
	BillTaskAcceptedCount int64 `json:"bill_task_accepted_count,omitempty" xml:"bill_task_accepted_count,omitempty"`
	// 其中已经接收到正确回执的单据数量是
	BillFileAcceptedCount int64 `json:"bill_file_accepted_count,omitempty" xml:"bill_file_accepted_count,omitempty"`
	// 其中单据生成包的数量
	BillTaskCount int64 `json:"bill_task_count,omitempty" xml:"bill_task_count,omitempty"`
	// 其中还没有接收回执的单据数量是
	BillFileUnAcceptedCount int64 `json:"bill_file_un_accepted_count,omitempty" xml:"bill_file_un_accepted_count,omitempty"`
	// 其中已经接收到正确回执的文件数量是
	RelationTaskAcceptedCount int64 `json:"relation_task_accepted_count,omitempty" xml:"relation_task_accepted_count,omitempty"`
	// 其中还没有接收回执的关联关系数量是
	RelationFileUnAcceptedCount int64 `json:"relation_file_un_accepted_count,omitempty" xml:"relation_file_un_accepted_count,omitempty"`
	// 其中已经接收到正确回执的关联关系数量是
	RelationFileAcceptedCount int64 `json:"relation_file_accepted_count,omitempty" xml:"relation_file_accepted_count,omitempty"`
	// 其中还没有接收回执的数量是
	RelationTaskUnAcceptedCount int64 `json:"relation_task_un_accepted_count,omitempty" xml:"relation_task_un_accepted_count,omitempty"`
	// 其中还没有接收回执的包数量是
	BillTaskUnAcceptedCount int64 `json:"bill_task_un_accepted_count,omitempty" xml:"bill_task_un_accepted_count,omitempty"`
}
