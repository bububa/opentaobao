package legalsuit

// SealResponseModel 结构体
type SealResponseModel struct {
	// 批次id
	BatchId string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 案件id
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 委托id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}
