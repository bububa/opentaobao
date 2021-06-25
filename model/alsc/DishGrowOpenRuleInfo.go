package alsc

// DishGrowOpenRuleInfo 
type DishGrowOpenRuleInfo struct {

    // 是否已删除
    Deleted   bool `json:"deleted,omitempty"`

    // 规则ID
    DishRuleId   string `json:"dish_rule_id,omitempty"`

    // 扩展信息
    ExtInfo   string `json:"ext_info,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 会员计划ID
    PlanId   string `json:"plan_id,omitempty"`

    // 菜品ID
    SkuId   string `json:"sku_id,omitempty"`

    // 升级到的等级id
    ToLevelId   string `json:"to_level_id,omitempty"`

    // 外部菜品ID
    OutSkuId   string `json:"out_sku_id,omitempty"`

}
