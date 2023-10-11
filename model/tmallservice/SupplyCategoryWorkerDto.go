package tmallservice

// SupplyCategoryWorkerDto 结构体
type SupplyCategoryWorkerDto struct {
	// 工人身份证号
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 工人手机号
	WorkerMobile string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
	// 工人id
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
}
