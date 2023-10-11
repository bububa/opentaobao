package simba

// TimeSpanQueryResVo 结构体
type TimeSpanQueryResVo struct {
	// 折扣时间段
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 折扣
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}
