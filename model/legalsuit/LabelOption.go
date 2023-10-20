package legalsuit

import (
	"sync"
)

// LabelOption 结构体
type LabelOption struct {
	// 子对象
	Children []LabelOption `json:"children,omitempty" xml:"children>label_option,omitempty"`
	// 案件类型中文
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 案件类型中文value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolLabelOption = sync.Pool{
	New: func() any {
		return new(LabelOption)
	},
}

// GetLabelOption() 从对象池中获取LabelOption
func GetLabelOption() *LabelOption {
	return poolLabelOption.Get().(*LabelOption)
}

// ReleaseLabelOption 释放LabelOption
func ReleaseLabelOption(v *LabelOption) {
	v.Children = v.Children[:0]
	v.Text = ""
	v.Value = ""
	poolLabelOption.Put(v)
}
