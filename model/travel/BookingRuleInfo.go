package travel

import (
	"sync"
)

// BookingRuleInfo 结构体
type BookingRuleInfo struct {
	// 规则说明，1500个字符
	RuleList []string `json:"rule_list,omitempty" xml:"rule_list>string,omitempty"`
	// 1500个字
	RuleDesc string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
	// fee_included：费用包含，跟团游必填； fee_excluded：费用不含，所有类目必填； order_info：预定须知； extra_cost：其他费用，预留；
	RuleType string `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
}

var poolBookingRuleInfo = sync.Pool{
	New: func() any {
		return new(BookingRuleInfo)
	},
}

// GetBookingRuleInfo() 从对象池中获取BookingRuleInfo
func GetBookingRuleInfo() *BookingRuleInfo {
	return poolBookingRuleInfo.Get().(*BookingRuleInfo)
}

// ReleaseBookingRuleInfo 释放BookingRuleInfo
func ReleaseBookingRuleInfo(v *BookingRuleInfo) {
	v.RuleList = v.RuleList[:0]
	v.RuleDesc = ""
	v.RuleType = ""
	poolBookingRuleInfo.Put(v)
}
