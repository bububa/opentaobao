package alihealthpw

// PendingListDto 结构体
type PendingListDto struct {
	// 申请状态
	ApplyStatus string `json:"apply_status,omitempty" xml:"apply_status,omitempty"`
	// 申请类型
	ApplyType string `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
	// 申请时间
	ApplyDate string `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 唯一编码
	ApplyUniqueCode string `json:"apply_unique_code,omitempty" xml:"apply_unique_code,omitempty"`
}
