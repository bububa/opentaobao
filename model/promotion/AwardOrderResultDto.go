package promotion

// AwardOrderResultDto 结构体
type AwardOrderResultDto struct {
	// 订单列表
	Orders []AwardOrder `json:"orders,omitempty" xml:"orders>award_order,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
