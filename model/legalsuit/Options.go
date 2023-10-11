package legalsuit

// Options 结构体
type Options struct {
	// 文件
	FileValues []FileValues `json:"file_values,omitempty" xml:"file_values>file_values,omitempty"`
	// 标识
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
