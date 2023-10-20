package promotion

import (
	"sync"
)

// ShowStrategyDto 结构体
type ShowStrategyDto struct {
	// 投放计划规则
	ShowRules []ShowRuleDto `json:"show_rules,omitempty" xml:"show_rules>show_rule_dto,omitempty"`
	// 投放计划维度核销数据
	ShowBenefitInstances []ShowBenefitInstanceDto `json:"show_benefit_instances,omitempty" xml:"show_benefit_instances>show_benefit_instance_dto,omitempty"`
	// 投放模式
	Mode string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 投放计划code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 投放计划安全码
	Asac string `json:"asac,omitempty" xml:"asac,omitempty"`
	// 是否投放计划所有规则通过
	AllRulePassed bool `json:"all_rule_passed,omitempty" xml:"all_rule_passed,omitempty"`
	// 算法容灾结果
	AlgorithmFailover bool `json:"algorithm_failover,omitempty" xml:"algorithm_failover,omitempty"`
}

var poolShowStrategyDto = sync.Pool{
	New: func() any {
		return new(ShowStrategyDto)
	},
}

// GetShowStrategyDto() 从对象池中获取ShowStrategyDto
func GetShowStrategyDto() *ShowStrategyDto {
	return poolShowStrategyDto.Get().(*ShowStrategyDto)
}

// ReleaseShowStrategyDto 释放ShowStrategyDto
func ReleaseShowStrategyDto(v *ShowStrategyDto) {
	v.ShowRules = v.ShowRules[:0]
	v.ShowBenefitInstances = v.ShowBenefitInstances[:0]
	v.Mode = ""
	v.Code = ""
	v.Asac = ""
	v.AllRulePassed = false
	v.AlgorithmFailover = false
	poolShowStrategyDto.Put(v)
}
