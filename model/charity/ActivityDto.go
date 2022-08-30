package charity

// ActivityDto 结构体
type ActivityDto struct {
	// 活动标题
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动摘要
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 活动内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}
