package alihealthpw

// PendingListResp 结构体
type PendingListResp struct {
	// 列表
	List []PendingListDto `json:"list,omitempty" xml:"list>pending_list_dto,omitempty"`
	// 申请类型
	ApplyType string `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 当前页
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
