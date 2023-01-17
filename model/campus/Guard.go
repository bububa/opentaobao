package campus

// Guard 结构体
type Guard struct {
	// 门禁点名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 门禁点id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
