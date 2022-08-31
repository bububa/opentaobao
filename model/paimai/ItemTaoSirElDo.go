package paimai

// ItemTaoSirElDo 结构体
type ItemTaoSirElDo struct {
	// 显示文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 0 - 类型为label元素，只用于展示，不用于组装value_name；&lt;br/&gt;1 - 类型为label元素，用于展示，用于组装value_name；&lt;br/&gt;2 - 类型为输入狂元素，主要用于卖家输入数据. 卖家填写完后需要重新设置该元素的文本数据；
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 是否只用于显示的文本元素
	IsShowLabel bool `json:"is_show_label,omitempty" xml:"is_show_label,omitempty"`
	// 是否文本元素，用于显示也用于组装value_name，比如：操作符
	IsLabel bool `json:"is_label,omitempty" xml:"is_label,omitempty"`
	// 是否输入框
	IsInput bool `json:"is_input,omitempty" xml:"is_input,omitempty"`
}
