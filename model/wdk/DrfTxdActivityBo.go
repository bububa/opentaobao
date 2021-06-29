package wdk

// DrfTxdActivityBo 
type DrfTxdActivityBo struct {
    // 换购机台
    ActivityChannel   string `json:"activity_channel,omitempty" xml:"activity_channel,omitempty"`
    // 封顶金额，单位：分
    CellingAmount   int64 `json:"celling_amount,omitempty" xml:"celling_amount,omitempty"`
    // 是否多阶梯可叠加，0--否，1--是
    IsMultiMix   int64 `json:"is_multi_mix,omitempty" xml:"is_multi_mix,omitempty"`
    // 是否单商品累计，0--否，1--是
    ItemOverlay   int64 `json:"item_overlay,omitempty" xml:"item_overlay,omitempty"`
    // 是否可贬值；0--否，1--是
    DiscountFeeModel   int64 `json:"discount_fee_model,omitempty" xml:"discount_fee_model,omitempty"`
    // 是否上不封顶；0--否，1--是
    EnableMultiple   int64 `json:"enable_multiple,omitempty" xml:"enable_multiple,omitempty"`
    // 商品池规则
    PoolRules   string `json:"pool_rules,omitempty" xml:"pool_rules,omitempty"`
    // 商品池数量
    PoolNum   int64 `json:"pool_num,omitempty" xml:"pool_num,omitempty"`
    // 商家人群编码
    MerchantCrowdCode   string `json:"merchant_crowd_code,omitempty" xml:"merchant_crowd_code,omitempty"`
    // 活动每日限购
    TotalDayLimit   int64 `json:"total_day_limit,omitempty" xml:"total_day_limit,omitempty"`
    // 用户每日限购
    UserDayLimit   int64 `json:"user_day_limit,omitempty" xml:"user_day_limit,omitempty"`
    // 活动总限购
    TotalLimit   int64 `json:"total_limit,omitempty" xml:"total_limit,omitempty"`
    // 用户数量总限购
    UserLimit   int64 `json:"user_limit,omitempty" xml:"user_limit,omitempty"`
    // 一口价【分】
    FixPrice   int64 `json:"fix_price,omitempty" xml:"fix_price,omitempty"`
    // 第N件活动规则
    CountAt   int64 `json:"count_at,omitempty" xml:"count_at,omitempty"`
    // 单品特价类型；1-一口价；2-打折；3-减钱
    DiscountType   int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
    // ?营销中台活动类型
    ActivityType   int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
    // 1--pos,2--App;1,2--pos&App
    Terminals   string `json:"terminals,omitempty" xml:"terminals,omitempty"`
    // 门店Id
    StoreIds   string `json:"store_ids,omitempty" xml:"store_ids,omitempty"`
    // 活动描述
    ActivityContent   string `json:"activity_content,omitempty" xml:"activity_content,omitempty"`
    // 活动名称
    ActivityName   string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
    // 中台活动Id（全局唯一）
    PromotionId   string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
    // 限购权重，实际限购=限购值/限购权重
    LimitWeight   int64 `json:"limit_weight,omitempty" xml:"limit_weight,omitempty"`
    // 更新时间
    UpdateTime   int64 `json:"update_time,omitempty" xml:"update_time,omitempty"`
    // 插入时间
    InsertTime   int64 `json:"insert_time,omitempty" xml:"insert_time,omitempty"`
    // 0--不可用；1--可用
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 活动结束时间
    EndDate   int64 `json:"end_date,omitempty" xml:"end_date,omitempty"`
    // 活动开始时间
    StartDate   int64 `json:"start_date,omitempty" xml:"start_date,omitempty"`
    // 商品池阶梯规则
    StairRules   string `json:"stair_rules,omitempty" xml:"stair_rules,omitempty"`
    // 是否叠加逻辑分组与阶梯条件；0-否，1-是
    IsCheckAllCond   int64 `json:"is_check_all_cond,omitempty" xml:"is_check_all_cond,omitempty"`
    // 淘鲜达活动Id
    TxdActivityId   int64 `json:"txd_activity_id,omitempty" xml:"txd_activity_id,omitempty"`
}
