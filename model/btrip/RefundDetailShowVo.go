package btrip

// RefundDetailShowVo 结构体
type RefundDetailShowVo struct {
	// 退改表达内容
	RefundSubItems []RefundSubItem `json:"refund_sub_items,omitempty" xml:"refund_sub_items>refund_sub_item,omitempty"`
	// 纯文字段落
	ExtraContents []RefundSubItem `json:"extra_contents,omitempty" xml:"extra_contents>refund_sub_item,omitempty"`
	// 行李表达内容
	BaggageSubItems []BaggageSubItem `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items>baggage_sub_item,omitempty"`
	// 表格标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 表格首行内容
	TableHead string `json:"table_head,omitempty" xml:"table_head,omitempty"`
	// 内容类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 排序使用
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
	// 提示
	Tips *Tips `json:"tips,omitempty" xml:"tips,omitempty"`
}
