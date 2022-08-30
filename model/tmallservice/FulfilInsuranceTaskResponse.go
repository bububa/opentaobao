package tmallservice

// FulfilInsuranceTaskResponse 结构体
type FulfilInsuranceTaskResponse struct {
	// 外部单号id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 履行单id
	FulfilTaskId int64 `json:"fulfil_task_id,omitempty" xml:"fulfil_task_id,omitempty"`
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}
