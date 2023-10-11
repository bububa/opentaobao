package wdk

// ShopInfoUpdateRequest 结构体
type ShopInfoUpdateRequest struct {
	// 营业开始时间(HH:mm)
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 营业结束时间(HH:mm)
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道
	ChannelSourceType int64 `json:"channel_source_type,omitempty" xml:"channel_source_type,omitempty"`
}
