package wdk

// ActivitySkuQueryDto 结构体
type ActivitySkuQueryDto struct {
	// 当前页码，从1开始
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
