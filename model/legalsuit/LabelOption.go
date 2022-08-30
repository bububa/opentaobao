package legalsuit

// LabelOption 结构体
type LabelOption struct {
	// 子对象
	Children []LabelOption `json:"children,omitempty" xml:"children>label_option,omitempty"`
	// 案件类型中文
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 案件类型中文value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
