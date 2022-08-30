package alihouse

// SyncHouseTradeItemDto 结构体
type SyncHouseTradeItemDto struct {
	// 1
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 1
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 1
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 1
	BillPayItemId int64 `json:"bill_pay_item_id,omitempty" xml:"bill_pay_item_id,omitempty"`
	// 1
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 1
	RelationItemBizType int64 `json:"relation_item_biz_type,omitempty" xml:"relation_item_biz_type,omitempty"`
}
