package icbuseller

// QueryTradeDto 结构体
type QueryTradeDto struct {
	// 交易单状态变更起始时间
	FireTimeStart string `json:"fire_time_start,omitempty" xml:"fire_time_start,omitempty"`
	// 服务code列表
	ServiceCode []string `json:"service_code,omitempty" xml:"service_code>string,omitempty"`
	// 起始值
	OffSet int64 `json:"off_set,omitempty" xml:"off_set,omitempty"`
	// 交易单创建结束时间
	CreateTimeEnd string `json:"create_time_end,omitempty" xml:"create_time_end,omitempty"`
	// 结束值
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 交易单创建起始时间
	CreateTimeStart string `json:"create_time_start,omitempty" xml:"create_time_start,omitempty"`
	// 是否展示
	IsDisplay bool `json:"is_display,omitempty" xml:"is_display,omitempty"`
	// 交易单状态变更结束时间
	FireTimeEnd string `json:"fire_time_end,omitempty" xml:"fire_time_end,omitempty"`
	// 买家aliid
	BuyerAliId int64 `json:"buyer_ali_id,omitempty" xml:"buyer_ali_id,omitempty"`
	// 交易id号列表
	TradeIds []int64 `json:"trade_ids,omitempty" xml:"trade_ids>int64,omitempty"`
	// 订单号列表
	OrderNos []string `json:"order_nos,omitempty" xml:"order_nos>string,omitempty"`
	// 状态列表
	Status []Null `json:"status,omitempty" xml:"status>null,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页显示数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
