package promotion

// CouponTemplateConditionConfig 结构体
type CouponTemplateConditionConfig struct {
	// 生效类目
	Categories []string `json:"categories,omitempty" xml:"categories>string,omitempty"`
	// 生效门店
	ShopIds []string `json:"shop_ids,omitempty" xml:"shop_ids>string,omitempty"`
	// 生效终端:  1.app 2.pos
	Terminals []string `json:"terminals,omitempty" xml:"terminals>string,omitempty"`
	// 排除的商品规则 1:专柜品
	ExcludeItemRules []string `json:"exclude_item_rules,omitempty" xml:"exclude_item_rules>string,omitempty"`
	// 生效商家类目
	MerchantCategories []string `json:"merchant_categories,omitempty" xml:"merchant_categories>string,omitempty"`
	// 门槛金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 门槛件数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 券的使用范围 COUPON_RANGE_STORE(1, "rangeShop", "店铺券"), COUPON_RANGE_ITEM(2, "rangeItem", "商品券"), COUPON_RANGE_CATEGORY(3, "rangeCategory", "品类券"), COUPON_RANGE_SELLER(307,"rangeSeller", "卖家券"),
	RangeType int64 `json:"range_type,omitempty" xml:"range_type,omitempty"`
	// 人群信息
	UserCrowdConfig *CouponTemplateUserCrowdConfig `json:"user_crowd_config,omitempty" xml:"user_crowd_config,omitempty"`
	// 单笔订单最大可用张数
	MaxUseCountPerOrder int64 `json:"max_use_count_per_order,omitempty" xml:"max_use_count_per_order,omitempty"`
	// 最大优惠金额
	MaxEffectAmount int64 `json:"max_effect_amount,omitempty" xml:"max_effect_amount,omitempty"`
	// 是否生效满元条件
	AmountAt bool `json:"amount_at,omitempty" xml:"amount_at,omitempty"`
	// 是否生效满件条件
	CountAt bool `json:"count_at,omitempty" xml:"count_at,omitempty"`
	// 是否限制本人使用
	OnlyMemberSelf bool `json:"only_member_self,omitempty" xml:"only_member_self,omitempty"`
}
