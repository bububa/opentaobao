package alihealth2

// Patient 结构体
type Patient struct {
	// 人群标签
	Labels []string `json:"labels,omitempty" xml:"labels>string,omitempty"`
	// 淘系用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用药人性别（0未知1男2女）
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 用药人体重-单位kg
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 用药人年龄-单位：天
	Age int64 `json:"age,omitempty" xml:"age,omitempty"`
	// 用药人身高-单位cm
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}
