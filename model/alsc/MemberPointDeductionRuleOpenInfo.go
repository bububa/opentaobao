package alsc

import (
	"sync"
)

// MemberPointDeductionRuleOpenInfo 结构体
type MemberPointDeductionRuleOpenInfo struct {
	// 会员等级ID
	LevelId string `json:"level_id,omitempty" xml:"level_id,omitempty"`
	// 消费金额
	ConsumeMoney int64 `json:"consume_money,omitempty" xml:"consume_money,omitempty"`
	// 抵扣积分
	DeductedPoint int64 `json:"deducted_point,omitempty" xml:"deducted_point,omitempty"`
	// 单次使用的抵扣上限
	MaxPoint int64 `json:"max_point,omitempty" xml:"max_point,omitempty"`
	// 是否允许该等级进行积分抵现
	Enable bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 单次使用是否有抵扣上限
	HasUpperLimit bool `json:"has_upper_limit,omitempty" xml:"has_upper_limit,omitempty"`
}

var poolMemberPointDeductionRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(MemberPointDeductionRuleOpenInfo)
	},
}

// GetMemberPointDeductionRuleOpenInfo() 从对象池中获取MemberPointDeductionRuleOpenInfo
func GetMemberPointDeductionRuleOpenInfo() *MemberPointDeductionRuleOpenInfo {
	return poolMemberPointDeductionRuleOpenInfo.Get().(*MemberPointDeductionRuleOpenInfo)
}

// ReleaseMemberPointDeductionRuleOpenInfo 释放MemberPointDeductionRuleOpenInfo
func ReleaseMemberPointDeductionRuleOpenInfo(v *MemberPointDeductionRuleOpenInfo) {
	v.LevelId = ""
	v.ConsumeMoney = 0
	v.DeductedPoint = 0
	v.MaxPoint = 0
	v.Enable = false
	v.HasUpperLimit = false
	poolMemberPointDeductionRuleOpenInfo.Put(v)
}
