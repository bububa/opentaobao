package flight

// DatePairDto 结构体
type DatePairDto struct {
	// 允许航班截止日期，无需传入时分秒
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 允许航班起始日期，无需传入时分秒
	Start string `json:"start,omitempty" xml:"start,omitempty"`
}
