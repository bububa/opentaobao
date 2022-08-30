package ascp

// PageData 结构体
type PageData struct {
	// 发货单列表
	List []ConsignOrder `json:"list,omitempty" xml:"list>consign_order,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页数量，不大于100
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
