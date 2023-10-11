package cainiaolocker

// Item 结构体
type Item struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
