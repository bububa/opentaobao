package btrip

// RefundChangeRuleItem 结构体
type RefundChangeRuleItem struct {
	// 说明文案
	ExtraContents []ExtraContentsBean `json:"extra_contents,omitempty" xml:"extra_contents>extra_contents_bean,omitempty"`
	// 与子内容列数保持一致
	SubTableHead []string `json:"sub_table_head,omitempty" xml:"sub_table_head>string,omitempty"`
	// 退改表达内容 �
	RefundSubItems []RefundChangeRuleSubItem `json:"refund_sub_items,omitempty" xml:"refund_sub_items>refund_change_rule_sub_item,omitempty"`
	// 表格首行内容(去程 北京-上海)
	TableHead string `json:"table_head,omitempty" xml:"table_head,omitempty"`
	// 表格标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 排序使用
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
	// 内容类型(0退票/1改期)
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
