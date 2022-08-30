package pur

// RelationTopDto 结构体
type RelationTopDto struct {
	// 拓展字段
	ExtStr string `json:"ext_str,omitempty" xml:"ext_str,omitempty"`
	// 饿了么行id
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 生成的发货行id
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
