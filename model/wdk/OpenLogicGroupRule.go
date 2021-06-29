package wdk

// OpenLogicGroupRule 
type OpenLogicGroupRule struct {
    // 整体优惠
    CoverAllDiscountRule   *CoverAllDiscountRule `json:"cover_all_discount_rule,omitempty" xml:"cover_all_discount_rule,omitempty"`
    // 是否为优惠作用分组
    IsEffectiveGroup   bool `json:"is_effective_group,omitempty" xml:"is_effective_group,omitempty"`
    // 逻辑分组上满N元条件，当阶梯的is_amount=true时生效，逻辑分组条件的优先级大于阶梯条件的优先级
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    // 逻辑分组上满N件条件，当阶梯的is_count=true时生效，逻辑分组条件的优先级大于阶梯条件的优先级
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // 换购N件
    CanExtraItemNum   int64 `json:"can_extra_item_num,omitempty" xml:"can_extra_item_num,omitempty"`
    // 1-普通分组，2-换购分组，3-买赠分组
    LogicGroupType   int64 `json:"logic_group_type,omitempty" xml:"logic_group_type,omitempty"`
    // 商家逻辑分组Id（单个活动不能重复）
    Number   int64 `json:"number,omitempty" xml:"number,omitempty"`
    // 换购分组排序，用于在app上展示的顺序
    ExchangeGroupOrder   int64 `json:"exchange_group_order,omitempty" xml:"exchange_group_order,omitempty"`
    // 换购分组名称
    ExchangeGroupName   string `json:"exchange_group_name,omitempty" xml:"exchange_group_name,omitempty"`
    // 分摊比例
    Ratio   int64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
}
