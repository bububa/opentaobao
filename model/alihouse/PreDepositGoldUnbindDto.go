package alihouse

// PreDepositGoldUnbindDto 结构体
type PreDepositGoldUnbindDto struct {
	// 外部活动id
	OuterSalesActivityId string `json:"outer_sales_activity_id,omitempty" xml:"outer_sales_activity_id,omitempty"`
	// 预存金淘宝商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
