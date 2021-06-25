package wdk

// ActivityRule 
type ActivityRule struct {

    // 封顶金额
    CeilingAmount   int64 `json:"ceiling_amount,omitempty"`

    // 多阶梯是否可叠加
    IsMultiMix   bool `json:"is_multi_mix,omitempty"`

    // 单商品是否累计
    ItemOverlay   bool `json:"item_overlay,omitempty"`

    // 是否单商品计数【NY使用】【废弃】
    IsAlone   bool `json:"is_alone,omitempty"`

    // 是否上不封顶
    EnableMultiple   bool `json:"enable_multiple,omitempty"`

    // 是否叠加计算逻辑分组与阶梯满元【件】条件
    IsCheckAllCond   bool `json:"is_check_all_cond,omitempty"`

    // 1-可贬值，0-不可贬值
    DiscountFeeMode   int64 `json:"discount_fee_mode,omitempty"`

}
