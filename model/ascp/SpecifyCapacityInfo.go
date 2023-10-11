package ascp

// SpecifyCapacityInfo 结构体
type SpecifyCapacityInfo struct {
	// 时间段产能
	SpecifyDateCapacity []CapacityInfo `json:"specify_date_capacity,omitempty" xml:"specify_date_capacity>capacity_info,omitempty"`
	// 指定日期，YYYY-MM-DD
	SpecifyDate string `json:"specify_date,omitempty" xml:"specify_date,omitempty"`
}
