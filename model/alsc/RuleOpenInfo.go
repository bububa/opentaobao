package alsc

// RuleOpenInfo 
/* model for simplify = false
type RuleOpenInfo struct {

    // 积分规则
    
    PointDeductionRule  *struct {
        PointDeductionRuleOpenInfo  *PointDeductionRuleOpenInfo `json:"point_deduction_rule_open_info,omitempty"`
    } `json:"point_deduction_rule,omitempty"`
    

    // 储值规则
    
    RechargeRuleOpenInfo  *struct {
        RechargeRuleOpenInfo  *RechargeRuleOpenInfo `json:"recharge_rule_open_info,omitempty"`
    } `json:"recharge_rule_open_info,omitempty"`
    

    // 券模板规则设置
    
    VoucherTemplateSettingOpenInfos  struct {
        VoucherTemplateSettingOpenInfo  []VoucherTemplateSettingOpenInfo `json:"voucher_template_setting_open_info,omitempty"`
    } `json:"voucher_template_setting_open_infos,omitempty"`
    

}
*/

// RuleOpenInfo 
type RuleOpenInfo struct {

    // 积分规则
    PointDeductionRule   *PointDeductionRuleOpenInfo `json:"point_deduction_rule,omitempty"`

    // 储值规则
    RechargeRuleOpenInfo   *RechargeRuleOpenInfo `json:"recharge_rule_open_info,omitempty"`

    // 券模板规则设置
    VoucherTemplateSettingOpenInfos   []VoucherTemplateSettingOpenInfo `json:"voucher_template_setting_open_infos,omitempty"`

}
