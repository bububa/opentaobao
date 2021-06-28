package wdk

// ItemPoolActivity 
type ItemPoolActivity struct {

    // 活动的梯度列表
    
    RuleStairs   []Rulestairs `json:"rule_stairs,omitempty" xml:"rule_stairs,omitempty"`
    

    // 通用限购信息，-1为不限制，默认为不限制
    
    LimitInfo   *LimitInfo `json:"limit_info,omitempty" xml:"limit_info,omitempty"`
    

    // 商家活动id
    
    OutActId   string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
    

    // 报名活动Id
    
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    

    // 活动名称
    
    ActivityName   string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
    

    // 活动描述
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // startTime
    
    StartTime   int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

    // endTime
    
    EndTime   int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // 优惠适用场景:APP|POS|POS+APP 分别对应的值为1|2|1,2
    
    Terminals   []int64 `json:"terminals,omitempty" xml:"terminals>int64,omitempty"`
    

    // 会员维度活动参与人群限制：-1:不限制 1:会员专享 2:非会员专享
    
    MemberLimit   int64 `json:"member_limit,omitempty" xml:"member_limit,omitempty"`
    

    // shopIds
    
    ShopIds   []string `json:"shop_ids,omitempty" xml:"shop_ids>string,omitempty"`
    

    // merchantCrowdCode
    
    MerchantCrowdCode   string `json:"merchant_crowd_code,omitempty" xml:"merchant_crowd_code,omitempty"`
    

    // txdCrowdCode
    
    TxdCrowdCode   string `json:"txd_crowd_code,omitempty" xml:"txd_crowd_code,omitempty"`
    

    // 逻辑分组规则
    
    LogicGroupRules   []OpenLogicGroupRule `json:"logic_group_rules,omitempty" xml:"logic_group_rules,omitempty"`
    

    // 活动优惠规则
    
    ActivityRule   *ActivityRule `json:"activity_rule,omitempty" xml:"activity_rule,omitempty"`
    

    // 周期优惠信息
    
    PeriodConfig   *PeriodConfig `json:"period_config,omitempty" xml:"period_config,omitempty"`
    

    // 是否是组合优惠
    
    IsComb   bool `json:"is_comb,omitempty" xml:"is_comb,omitempty"`
    

    // 活动优先级，值越大表示优先级越高，必须大于0
    
    PriorityValue   int64 `json:"priority_value,omitempty" xml:"priority_value,omitempty"`
    

    // 商品池是否排除特价
    
    ExcludeSingle   bool `json:"exclude_single,omitempty" xml:"exclude_single,omitempty"`
    

    // 是否是类目优惠
    
    IsCategory   bool `json:"is_category,omitempty" xml:"is_category,omitempty"`
    

}
