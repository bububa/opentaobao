package feedflow

// TimeSpanDto 结构体
type TimeSpanDto struct {
	// 时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 折扣率
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}
