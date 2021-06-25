package alsc

// OptPlanOpenInfo 
type OptPlanOpenInfo struct {

    // 运营计划id
    OptPlanId   string `json:"opt_plan_id,omitempty"`

    // 运营计划名称
    OptPlanName   string `json:"opt_plan_name,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 运营计划类型
    OptPlanType   int64 `json:"opt_plan_type,omitempty"`

    // 储值规则Id
    RechargeRuleId   string `json:"recharge_rule_id,omitempty"`

    // 储值规则名称
    RechargeRuleName   string `json:"recharge_rule_name,omitempty"`

    // 积分规则Id
    PointRuleId   string `json:"point_rule_id,omitempty"`

    // 积分规则名称
    PointRuleName   string `json:"point_rule_name,omitempty"`

    // 适用门店数
    ApplyShopCount   int64 `json:"apply_shop_count,omitempty"`

    // 门店组信息
    ShopGroupInfo   *ShopGroupOpenInfo `json:"shop_group_info,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 创建者
    CreateBy   string `json:"create_by,omitempty"`

    // 更新者
    UpdateBy   string `json:"update_by,omitempty"`

    // 逻辑删除标志
    Deleted   bool `json:"deleted,omitempty"`

}
