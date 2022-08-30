package alihouse

// SyncHouseStatusDto 结构体
type SyncHouseStatusDto struct {
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 委托商品id
	EntrustItemId int64 `json:"entrust_item_id,omitempty" xml:"entrust_item_id,omitempty"`
	// 上下架状态
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 1
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 1
	BillPayItemId int64 `json:"bill_pay_item_id,omitempty" xml:"bill_pay_item_id,omitempty"`
}
