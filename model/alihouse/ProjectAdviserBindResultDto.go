package alihouse

// ProjectAdviserBindResultDto 结构体
type ProjectAdviserBindResultDto struct {
	// 外部顾问id
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 外部楼盘id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
