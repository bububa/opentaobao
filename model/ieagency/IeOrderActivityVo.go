package ieagency

// IeOrderActivityVo 结构体
type IeOrderActivityVo struct {
	// activityName
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// activityPrice
	ActivityPrice int64 `json:"activity_price,omitempty" xml:"activity_price,omitempty"`
}
