package ascp

// Period 结构体
type Period struct {
	// 开始时间
	BeginTime int64 `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 结束时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}
