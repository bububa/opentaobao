package tmallnr

// NrZqsPlanDetailInfoDto 结构体
type NrZqsPlanDetailInfoDto struct {
	// 配送期号从1开始，一直到N
	SequenceNo int64 `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
	// 计划配送时间
	PlanDate string `json:"plan_date,omitempty" xml:"plan_date,omitempty"`
}
