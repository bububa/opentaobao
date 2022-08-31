package btrip

// BaggageItem 结构体
type BaggageItem struct {
	// 行李额子内容
	BaggageSubItems []BaggageSubItem `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items>baggage_sub_item,omitempty"`
	// 表格首行内容(去程 北京-上海)
	TableHead string `json:"table_head,omitempty" xml:"table_head,omitempty"`
	// 表格标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 排序使用
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
	// 内容类型(0退票/1改期)
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 提示
	Tips *BaggageTip `json:"tips,omitempty" xml:"tips,omitempty"`
}
