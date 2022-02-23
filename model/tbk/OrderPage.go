package tbk

// OrderPage 结构体
type OrderPage struct {
	// PublisherOrderDto
	Results []PublisherOrderDto `json:"results,omitempty" xml:"results>publisher_order_dto,omitempty"`
	// 位点字段，由调用方原样传递
	PositionIndex string `json:"position_index,omitempty" xml:"position_index,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否还有上一页
	HasPre bool `json:"has_pre,omitempty" xml:"has_pre,omitempty"`
	// 是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
