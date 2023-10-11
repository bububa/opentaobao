package axindata

// FscTravellerDefineApiDto 结构体
type FscTravellerDefineApiDto struct {
	// 出行人类型 ADULT:成人 CHILD:儿童
	TravellerType string `json:"traveller_type,omitempty" xml:"traveller_type,omitempty"`
	// 出行人类型说明
	TravellerDesc string `json:"traveller_desc,omitempty" xml:"traveller_desc,omitempty"`
	// 起始年龄，包含
	StartAge int64 `json:"start_age,omitempty" xml:"start_age,omitempty"`
	// 截止年龄，包含
	EndAge int64 `json:"end_age,omitempty" xml:"end_age,omitempty"`
}
