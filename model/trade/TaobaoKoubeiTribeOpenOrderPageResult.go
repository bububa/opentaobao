package trade

// TaobaoKoubeiTribeOpenOrderPageResult 结构体
type TaobaoKoubeiTribeOpenOrderPageResult struct {
	// 订单信息结果
	Data *OrderInfoResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误提示
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// request唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
}
