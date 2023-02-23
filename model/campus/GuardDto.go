package campus

// GuardDto 结构体
type GuardDto struct {
	// 门禁点名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 门禁点ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
