package logistic

// Capacities 结构体
type Capacities struct {
	// 门店编码，对应大润发deliveryDockCode
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 配送中骑手
	DeliveryKnightAmount int64 `json:"delivery_knight_amount,omitempty" xml:"delivery_knight_amount,omitempty"`
	// 小休骑手
	RestKnightAmount int64 `json:"rest_knight_amount,omitempty" xml:"rest_knight_amount,omitempty"`
	// 活跃骑手
	ActiveKnightAmount int64 `json:"active_knight_amount,omitempty" xml:"active_knight_amount,omitempty"`
	// 下班骑手
	OffWorkKnightAmount int64 `json:"off_work_knight_amount,omitempty" xml:"off_work_knight_amount,omitempty"`
	// 到店骑手
	ArrivalKnightAmount int64 `json:"arrival_knight_amount,omitempty" xml:"arrival_knight_amount,omitempty"`
	// 归途	骑手
	BackKnightAmount int64 `json:"back_knight_amount,omitempty" xml:"back_knight_amount,omitempty"`
	// 上班骑手数
	WorkKnightAmount int64 `json:"work_knight_amount,omitempty" xml:"work_knight_amount,omitempty"`
}
