package tmallservice

// SupplyLevelWorkerDto 结构体
type SupplyLevelWorkerDto struct {
	// serviceCode
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 工人身份证号
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 技能等级(中级、初级、试用、高级)
	WorkerLevel string `json:"worker_level,omitempty" xml:"worker_level,omitempty"`
	// 工人手机号
	WorkerMobile string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
	// 工人id
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
}
