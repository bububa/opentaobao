package alsc

// RechargeAccountOpenInfo 
/* model for simplify = false
type RechargeAccountOpenInfo struct {

    // 账户id
    
    AccountId   string `json:"account_id,omitempty"`
    

    // crm卡实例id
    
    CardId   string `json:"card_id,omitempty"`
    

    // 账户所属顾客id，如果是会员卡，不为空
    
    CustomerId   string `json:"customer_id,omitempty"`
    

    // 是否删除
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 账户赠储余额
    
    GiftValue   int64 `json:"gift_value,omitempty"`
    

    // 累计账户赠储余额
    
    GiftValueTotal   int64 `json:"gift_value_total,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 更新时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 关联的运营方案id
    
    OptPlanId   string `json:"opt_plan_id,omitempty"`
    

    // 账户预储余额
    
    PreValue   int64 `json:"pre_value,omitempty"`
    

    // 累计账户预储余额
    
    PreValueTotal   int64 `json:"pre_value_total,omitempty"`
    

    // 账户实储余额
    
    RealValue   int64 `json:"real_value,omitempty"`
    

    // 累计账户实储余额
    
    RealValueTotal   int64 `json:"real_value_total,omitempty"`
    

    // 可用余额
    
    UsableValue   int64 `json:"usable_value,omitempty"`
    

}
*/

// RechargeAccountOpenInfo 
type RechargeAccountOpenInfo struct {

    // 账户id
    AccountId   string `json:"account_id,omitempty"`

    // crm卡实例id
    CardId   string `json:"card_id,omitempty"`

    // 账户所属顾客id，如果是会员卡，不为空
    CustomerId   string `json:"customer_id,omitempty"`

    // 是否删除
    Deleted   bool `json:"deleted,omitempty"`

    // 账户赠储余额
    GiftValue   int64 `json:"gift_value,omitempty"`

    // 累计账户赠储余额
    GiftValueTotal   int64 `json:"gift_value_total,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 关联的运营方案id
    OptPlanId   string `json:"opt_plan_id,omitempty"`

    // 账户预储余额
    PreValue   int64 `json:"pre_value,omitempty"`

    // 累计账户预储余额
    PreValueTotal   int64 `json:"pre_value_total,omitempty"`

    // 账户实储余额
    RealValue   int64 `json:"real_value,omitempty"`

    // 累计账户实储余额
    RealValueTotal   int64 `json:"real_value_total,omitempty"`

    // 可用余额
    UsableValue   int64 `json:"usable_value,omitempty"`

}
