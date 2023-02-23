package axintrade

// ArrivalTimeDto 结构体
type ArrivalTimeDto struct {
	// 小时房可入住的最早时间
	CheckInStart string `json:"check_in_start,omitempty" xml:"check_in_start,omitempty"`
	// 小时房最晚离店时间
	LastLeaveTime string `json:"last_leave_time,omitempty" xml:"last_leave_time,omitempty"`
	// 连住时长
	Hourage int64 `json:"hourage,omitempty" xml:"hourage,omitempty"`
}
