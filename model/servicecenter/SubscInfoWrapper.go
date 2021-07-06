package servicecenter

// SubscInfoWrapper 结构体
type SubscInfoWrapper struct {
	// 需求订购信息
	SubscInfoList []SubscInfo `json:"subsc_info_list,omitempty" xml:"subsc_info_list>subsc_info,omitempty"`
	// 总量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
