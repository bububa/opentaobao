package scbp

// TagGroup 结构体
type TagGroup struct {
	// 分组名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 关键词数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 分组ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
