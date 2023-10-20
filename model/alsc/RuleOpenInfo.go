package alsc

import (
	"sync"
)

// RuleOpenInfo 结构体
type RuleOpenInfo struct {
	// 券模板规则设置
	VoucherTemplateSettingOpenInfos []VoucherTemplateSettingOpenInfo `json:"voucher_template_setting_open_infos,omitempty" xml:"voucher_template_setting_open_infos>voucher_template_setting_open_info,omitempty"`
	// 积分规则
	PointDeductionRule *PointDeductionRuleOpenInfo `json:"point_deduction_rule,omitempty" xml:"point_deduction_rule,omitempty"`
	// 储值规则
	RechargeRuleOpenInfo *RechargeRuleOpenInfo `json:"recharge_rule_open_info,omitempty" xml:"recharge_rule_open_info,omitempty"`
}

var poolRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(RuleOpenInfo)
	},
}

// GetRuleOpenInfo() 从对象池中获取RuleOpenInfo
func GetRuleOpenInfo() *RuleOpenInfo {
	return poolRuleOpenInfo.Get().(*RuleOpenInfo)
}

// ReleaseRuleOpenInfo 释放RuleOpenInfo
func ReleaseRuleOpenInfo(v *RuleOpenInfo) {
	v.VoucherTemplateSettingOpenInfos = v.VoucherTemplateSettingOpenInfos[:0]
	v.PointDeductionRule = nil
	v.RechargeRuleOpenInfo = nil
	poolRuleOpenInfo.Put(v)
}
