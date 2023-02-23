package tbtrade

// SortInfo 结构体
type SortInfo struct {
	// daySortVal
	DaySortVal int64 `json:"day_sort_val,omitempty" xml:"day_sort_val,omitempty"`
	// hourSortVal
	HourSortVal int64 `json:"hour_sort_val,omitempty" xml:"hour_sort_val,omitempty"`
}
