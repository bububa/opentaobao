package qimen

// EntryOrderQueryResponse 结构体
type EntryOrderQueryResponse struct {
	// 入库单单据信息
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// orderLines总条数
	TotalLines int64 `json:"totalLines,omitempty" xml:"totalLines,omitempty"`
	// 入库单信息
	EntryOrder *EntryOrder `json:"entryOrder,omitempty" xml:"entryOrder,omitempty"`
}
