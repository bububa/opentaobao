package qimen

// EntryOrderCreateRequest 结构体
type EntryOrderCreateRequest struct {
	// 入库单详情
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 入库单信息
	EntryOrder *EntryOrder `json:"entryOrder,omitempty" xml:"entryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenEntryorderCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
