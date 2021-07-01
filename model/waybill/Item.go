package waybill

// Item 结构体
type Item struct {
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
