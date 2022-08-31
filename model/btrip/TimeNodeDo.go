package btrip

// TimeNodeDo 结构体
type TimeNodeDo struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间类型
	TimeType string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	// 时间
	TimeStamp int64 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// 费用
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 费率
	CostPercent int64 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
}
