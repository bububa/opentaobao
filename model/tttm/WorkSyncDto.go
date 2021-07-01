package tttm

// WorkSyncDto 结构体
type WorkSyncDto struct {
	// 工单ID
	WorkId string `json:"work_id,omitempty" xml:"work_id,omitempty"`
	// 货品
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 计划产量
	PlanNum string `json:"plan_num,omitempty" xml:"plan_num,omitempty"`
	// 生产节点
	ProduceLinks []string `json:"produce_links,omitempty" xml:"produce_links>string,omitempty"`
	// 生产状态
	ProduceStatus int64 `json:"produce_status,omitempty" xml:"produce_status,omitempty"`
	// 工单开始
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 工单结束
	FinishTime string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
	// 生产进度
	SyncProduceDTOs []ProduceSyncDto `json:"sync_produce_d_t_os,omitempty" xml:"sync_produce_d_t_os>produce_sync_dto,omitempty"`
}
