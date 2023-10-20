package qimen

// TaobaoqimeninventorybatchqueryResponse 结构体
type TaobaoqimeninventorybatchqueryResponse struct {
	// success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总行数，必填
	TotalCount int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 明细列表
	Items *Items `json:"items,omitempty" xml:"items,omitempty"`
}
