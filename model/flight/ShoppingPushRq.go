package flight

// ShoppingPushRq 结构体
type ShoppingPushRq struct {
	// 去程日期，格式为 YYYYMMDD ；如果为多程，20180729,20180804方式传输数据
	FromDate string `json:"from_date,omitempty" xml:"from_date,omitempty"`
	// 回程日期，格式为 YYYYMMDD（不传此参数或者留空表示单程/多程）
	RetDate string `json:"ret_date,omitempty" xml:"ret_date,omitempty"`
	// 目的地城市 、到达城市IATA 三字码代码 ；多程为空
	ToCity string `json:"to_city,omitempty" xml:"to_city,omitempty"`
	// 政策详情，同大卖家API搜索结果返回的json字符串格式；
	SearchRs string `json:"search_rs,omitempty" xml:"search_rs,omitempty"`
	// 出发地 IATA 三字码代码; 如果为多程,最多三程六段，按照 PEK/HKG,HKG/SHA 格式请求
	FromCity string `json:"from_city,omitempty" xml:"from_city,omitempty"`
	// cid
	Cid string `json:"cid,omitempty" xml:"cid,omitempty"`
	// 渠道id
	ChannelId int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 行程类型，1：单程；2：往返；5:  多程
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}
