package alsc

// QueryCardOpenReq 结构体
type QueryCardOpenReq struct {
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卡实例id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 外部品牌id和品牌id必传1
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 外部门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 物理卡号
	PhysicalCardId string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
	// saas门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 查询关联资产账户
	AccountRequired bool `json:"account_required,omitempty" xml:"account_required,omitempty"`
	// 查询关联物理卡
	PhysicalCardRequired bool `json:"physical_card_required,omitempty" xml:"physical_card_required,omitempty"`
}
