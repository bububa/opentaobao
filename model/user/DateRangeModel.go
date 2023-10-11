package user

// DateRangeModel 结构体
type DateRangeModel struct {
	// 数据有效开始日期
	BeginDate int64 `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// 数据有效结束日期
	EndDate int64 `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 81或82
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 数据可用
	Readyable bool `json:"readyable,omitempty" xml:"readyable,omitempty"`
}
