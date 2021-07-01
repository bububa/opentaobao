package lstwarehouse

// WarehouseQueryParam 结构体
type WarehouseQueryParam struct {
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页最大记录数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}
