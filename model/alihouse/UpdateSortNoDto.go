package alihouse

// UpdateSortNoDto 结构体
type UpdateSortNoDto struct {
	// 楼盘外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 排序值 大于等于10的数字
	SortNo int64 `json:"sort_no,omitempty" xml:"sort_no,omitempty"`
}
