package alsc

import (
	"sync"
)

// PointClearRuleOpenInfo 结构体
type PointClearRuleOpenInfo struct {
	// 积分清零规则类型
	PointClearRuleType string `json:"point_clear_rule_type,omitempty" xml:"point_clear_rule_type,omitempty"`
	// 多少天之后清零
	Days int64 `json:"days,omitempty" xml:"days,omitempty"`
}

var poolPointClearRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(PointClearRuleOpenInfo)
	},
}

// GetPointClearRuleOpenInfo() 从对象池中获取PointClearRuleOpenInfo
func GetPointClearRuleOpenInfo() *PointClearRuleOpenInfo {
	return poolPointClearRuleOpenInfo.Get().(*PointClearRuleOpenInfo)
}

// ReleasePointClearRuleOpenInfo 释放PointClearRuleOpenInfo
func ReleasePointClearRuleOpenInfo(v *PointClearRuleOpenInfo) {
	v.PointClearRuleType = ""
	v.Days = 0
	poolPointClearRuleOpenInfo.Put(v)
}
