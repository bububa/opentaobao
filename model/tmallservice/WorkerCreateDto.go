package tmallservice

// WorkerCreateDto 结构体
type WorkerCreateDto struct {
	// 工人ID
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
}
