package alsc

// PointRuleOpenInfo 
type PointRuleOpenInfo struct {

    // 创建者
    CreateBy   string `json:"create_by,omitempty"`

    // 逻辑删除标志
    Deleted   bool `json:"deleted,omitempty"`

    // 扩展信息
    ExtInfo   *Extinfo `json:"ext_info,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 规则名称
    Name   string `json:"name,omitempty"`

    // 积分获取规则
    PointAdditionRule   *PointAdditionRuleOpenInfo `json:"point_addition_rule,omitempty"`

    // 积分清零规则
    PointClearRule   *PointClearRuleOpenInfo `json:"point_clear_rule,omitempty"`

    // 积分扣减规则
    PointDeductionRule   *PointDeductionRuleOpenInfo `json:"point_deduction_rule,omitempty"`

    // 规则的业务ID
    RuleId   string `json:"rule_id,omitempty"`

    // 更新者
    UpdateBy   string `json:"update_by,omitempty"`

    // 更新者姓名
    UpdateByName   string `json:"update_by_name,omitempty"`

    // 创建者姓名
    CreateByName   string `json:"create_by_name,omitempty"`

}
