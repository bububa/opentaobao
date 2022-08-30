package scbp

// TagDefineDto 结构体
type TagDefineDto struct {
	// 标签值
	OptionValue string `json:"option_value,omitempty" xml:"option_value,omitempty"`
	// 标签名（标签描述为空时，取标签名）
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签描述（标签名为空时，取标签描述）
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 层级（0,1,2）
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}
