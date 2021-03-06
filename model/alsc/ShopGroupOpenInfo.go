package alsc

// ShopGroupOpenInfo 结构体
type ShopGroupOpenInfo struct {
	// 门店id
	ShopIds []string `json:"shop_ids,omitempty" xml:"shop_ids>string,omitempty"`
	// 外部门店ID
	OutShopIds []string `json:"out_shop_ids,omitempty" xml:"out_shop_ids>string,omitempty"`
	// 门店组id
	ShopGroupId string `json:"shop_group_id,omitempty" xml:"shop_group_id,omitempty"`
	// 运营计划id
	OptPlanId string `json:"opt_plan_id,omitempty" xml:"opt_plan_id,omitempty"`
	// 门店组类型
	ShopGroupType int64 `json:"shop_group_type,omitempty" xml:"shop_group_type,omitempty"`
}
