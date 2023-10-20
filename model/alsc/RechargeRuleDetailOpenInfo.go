package alsc

import (
	"sync"
)

// RechargeRuleDetailOpenInfo 结构体
type RechargeRuleDetailOpenInfo struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 赠送积分
	GiftPoint int64 `json:"gift_point,omitempty" xml:"gift_point,omitempty"`
	// 赠送金额
	GiftValue int64 `json:"gift_value,omitempty" xml:"gift_value,omitempty"`
	// 储值金额
	RealValue int64 `json:"real_value,omitempty" xml:"real_value,omitempty"`
}

var poolRechargeRuleDetailOpenInfo = sync.Pool{
	New: func() any {
		return new(RechargeRuleDetailOpenInfo)
	},
}

// GetRechargeRuleDetailOpenInfo() 从对象池中获取RechargeRuleDetailOpenInfo
func GetRechargeRuleDetailOpenInfo() *RechargeRuleDetailOpenInfo {
	return poolRechargeRuleDetailOpenInfo.Get().(*RechargeRuleDetailOpenInfo)
}

// ReleaseRechargeRuleDetailOpenInfo 释放RechargeRuleDetailOpenInfo
func ReleaseRechargeRuleDetailOpenInfo(v *RechargeRuleDetailOpenInfo) {
	v.Remark = ""
	v.GiftPoint = 0
	v.GiftValue = 0
	v.RealValue = 0
	poolRechargeRuleDetailOpenInfo.Put(v)
}
