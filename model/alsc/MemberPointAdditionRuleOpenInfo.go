package alsc

// MemberPointAdditionRuleOpenInfo 
/* model for simplify = false
type MemberPointAdditionRuleOpenInfo struct {

    // 允许积分商品的范围
    
    AllowProductType   string `json:"allow_product_type,omitempty"`
    

    // 是否允许储值积分
    
    AllowRecharge   bool `json:"allow_recharge,omitempty"`
    

    // 消费的现金，单位为元
    
    ConsumeMoney   int64 `json:"consume_money,omitempty"`
    

    // 是否允许获取积分
    
    Enable   bool `json:"enable,omitempty"`
    

    // 会员等级ID
    
    LevelId   string `json:"level_id,omitempty"`
    

    // 奖励的积分
    
    RewardPoint   int64 `json:"reward_point,omitempty"`
    

    // 商品SKU_ID列表
    
    SkuIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sku_ids,omitempty"`
    

}
*/

// MemberPointAdditionRuleOpenInfo 
type MemberPointAdditionRuleOpenInfo struct {

    // 允许积分商品的范围
    AllowProductType   string `json:"allow_product_type,omitempty"`

    // 是否允许储值积分
    AllowRecharge   bool `json:"allow_recharge,omitempty"`

    // 消费的现金，单位为元
    ConsumeMoney   int64 `json:"consume_money,omitempty"`

    // 是否允许获取积分
    Enable   bool `json:"enable,omitempty"`

    // 会员等级ID
    LevelId   string `json:"level_id,omitempty"`

    // 奖励的积分
    RewardPoint   int64 `json:"reward_point,omitempty"`

    // 商品SKU_ID列表
    SkuIds   []string `json:"sku_ids,omitempty"`

}
