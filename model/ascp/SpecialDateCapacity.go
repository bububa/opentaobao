package ascp

// SpecialDateCapacity 结构体
type SpecialDateCapacity struct {
	// 日期（年月日），例如20230506
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 指定日期下产能（≥0）
	Capacity string `json:"capacity,omitempty" xml:"capacity,omitempty"`
}
