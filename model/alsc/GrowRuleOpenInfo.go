package alsc

// GrowRuleOpenInfo 
/* model for simplify = false
type GrowRuleOpenInfo struct {

    // 普通消费能否获取成长值
    
    CommonConsumeGrowSupport   bool `json:"common_consume_grow_support,omitempty"`
    

    // 退款是否扣除成长值
    
    DecreaseSupport   bool `json:"decrease_support,omitempty"`
    

    // 是否已删除
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 扩展信息
    
    ExtInfo   string `json:"ext_info,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 规则id
    
    GrowRuleId   string `json:"grow_rule_id,omitempty"`
    

    // 是否支持成长值获取(总开关)
    
    GrowSupport   bool `json:"grow_support,omitempty"`
    

    // 储值消费能否获取成长值
    
    RechargeConsumeGrowSupport   bool `json:"recharge_consume_grow_support,omitempty"`
    

    // 储值能否获取成长值
    
    RechargeGrowSupport   bool `json:"recharge_grow_support,omitempty"`
    

    // 创建人
    
    CreateBy   string `json:"create_by,omitempty"`
    

    // 不同等级消费获取成长值规则模型
    
    LevelConsumeGrowRuleOpenInfoList  struct {
        LevelConsumeGrowRuleOpenInfo  []LevelConsumeGrowRuleOpenInfo `json:"level_consume_grow_rule_open_info,omitempty"`
    } `json:"level_consume_grow_rule_open_info_list,omitempty"`
    

    // 更新人
    
    UpdateBy   string `json:"update_by,omitempty"`
    

}
*/

// GrowRuleOpenInfo 
type GrowRuleOpenInfo struct {

    // 普通消费能否获取成长值
    CommonConsumeGrowSupport   bool `json:"common_consume_grow_support,omitempty"`

    // 退款是否扣除成长值
    DecreaseSupport   bool `json:"decrease_support,omitempty"`

    // 是否已删除
    Deleted   bool `json:"deleted,omitempty"`

    // 扩展信息
    ExtInfo   string `json:"ext_info,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 规则id
    GrowRuleId   string `json:"grow_rule_id,omitempty"`

    // 是否支持成长值获取(总开关)
    GrowSupport   bool `json:"grow_support,omitempty"`

    // 储值消费能否获取成长值
    RechargeConsumeGrowSupport   bool `json:"recharge_consume_grow_support,omitempty"`

    // 储值能否获取成长值
    RechargeGrowSupport   bool `json:"recharge_grow_support,omitempty"`

    // 创建人
    CreateBy   string `json:"create_by,omitempty"`

    // 不同等级消费获取成长值规则模型
    LevelConsumeGrowRuleOpenInfoList   []LevelConsumeGrowRuleOpenInfo `json:"level_consume_grow_rule_open_info_list,omitempty"`

    // 更新人
    UpdateBy   string `json:"update_by,omitempty"`

}
