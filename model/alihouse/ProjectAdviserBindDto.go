package alihouse

// ProjectAdviserBindDto 结构体
type ProjectAdviserBindDto struct {
	// 外部顾问id
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 外部楼盘id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}
