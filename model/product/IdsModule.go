package product

// IdsModule 结构体
type IdsModule struct {
	// 宝贝描述规范化模块id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 宝贝描述规范化模块名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 0为自动打标；<br/>1为人工打标；
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
