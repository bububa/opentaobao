package alihouse

// ProjectCardDataDto 结构体
type ProjectCardDataDto struct {
	// 楼盘外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 卡片占位符
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}
