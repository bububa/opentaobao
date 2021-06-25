package promotion

// ExpressionRuleDto 
type ExpressionRuleDto struct {

    // 规则ID
    Id   int64 `json:"id,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 规则名称
    Name   string `json:"name,omitempty"`

    // 规则主体id
    MasterId   int64 `json:"master_id,omitempty"`

    // 规则主体类型（1活动2奖品3方案）
    MasterType   int64 `json:"master_type,omitempty"`

    // 规则CODE
    RuleCode   string `json:"rule_code,omitempty"`

    // 规则参数
    Params   string `json:"params,omitempty"`

}
