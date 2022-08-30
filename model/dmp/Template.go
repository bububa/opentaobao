package dmp

// Template 结构体
type Template struct {
	// 模版扩展json信息，usageIndex: 使用热度，mark: 打标文案
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 模版描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 模版推荐理由
	Highlight string `json:"highlight,omitempty" xml:"highlight,omitempty"`
	// 模版名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 模版id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
