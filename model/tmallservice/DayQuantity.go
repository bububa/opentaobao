package tmallservice

// DayQuantity 结构体
type DayQuantity struct {
	// 必须是YYYY-MM-DD格式
	Day string `json:"day,omitempty" xml:"day,omitempty"`
	// 表示具体的某一天内的一个时间段。注意：如果配置了改值，则这里容量调整的值则是2022-07-01 08:00-09:00的容量，不配置则是调整07-01这一天的容量
	TimeRange string `json:"time_range,omitempty" xml:"time_range,omitempty"`
	// 容量值为1
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 非必填，
	Available bool `json:"available,omitempty" xml:"available,omitempty"`
}
