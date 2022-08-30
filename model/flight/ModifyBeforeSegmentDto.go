package flight

// ModifyBeforeSegmentDto 结构体
type ModifyBeforeSegmentDto struct {
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 起飞城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
}
