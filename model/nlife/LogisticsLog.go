package nlife

// LogisticsLog 结构体
type LogisticsLog struct {
	// time
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 内容
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
}
