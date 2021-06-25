package promotion

// CouponTemplateConditionConfig 
type CouponTemplateConditionConfig struct {

    // 门槛金额
    Amount   int64 `json:"amount,omitempty"`

    // 是否生效满元条件
    AmountAt   bool `json:"amount_at,omitempty"`

    // 生效类目
    Categories   []String `json:"categories,omitempty"`

    // 门槛件数
    Count   int64 `json:"count,omitempty"`

    // 是否生效满件条件
    CountAt   bool `json:"count_at,omitempty"`

    // 是否限制本人使用
    OnlyMemberSelf   bool `json:"only_member_self,omitempty"`

    // 券的使用范围 COUPON_RANGE_STORE(1, "rangeShop", "店铺券"), COUPON_RANGE_ITEM(2, "rangeItem", "商品券"), COUPON_RANGE_CATEGORY(3, "rangeCategory", "品类券"), COUPON_RANGE_SELLER(307,"rangeSeller", "卖家券"),
    RangeType   int64 `json:"range_type,omitempty"`

    // 生效门店
    ShopIds   []String `json:"shop_ids,omitempty"`

    // 生效终端:  1.app 2.pos
    Terminals   []Number `json:"terminals,omitempty"`

    // 人群信息
    UserCrowdConfig   *CouponTemplateUserCrowdConfig `json:"user_crowd_config,omitempty"`

    // 单笔订单最大可用张数
    MaxUseCountPerOrder   int64 `json:"max_use_count_per_order,omitempty"`

    // 最大优惠金额
    MaxEffectAmount   int64 `json:"max_effect_amount,omitempty"`

    // 排除的商品规则 1:专柜品
    ExcludeItemRules   []Number `json:"exclude_item_rules,omitempty"`

    // 生效商家类目
    MerchantCategories   []String `json:"merchant_categories,omitempty"`

}
