package travel

import (
	"sync"
)

// PontusTravelBookingRuleInfo 结构体
type PontusTravelBookingRuleInfo struct {
	// 分点组织的规则说明
	RuleList []string `json:"rule_list,omitempty" xml:"rule_list>string,omitempty"`
	// 规则说明，1500个字符
	RuleDesc string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
	// Fee_Included：费用包含，跟团游必填； Fee_Excluded：费用不含，所有类目必填； Order_Info：预定须知； Extra_Cost：其他费用，预留；     Shopping：购物说明，预留
	RuleType string `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
}

var poolPontusTravelBookingRuleInfo = sync.Pool{
	New: func() any {
		return new(PontusTravelBookingRuleInfo)
	},
}

// GetPontusTravelBookingRuleInfo() 从对象池中获取PontusTravelBookingRuleInfo
func GetPontusTravelBookingRuleInfo() *PontusTravelBookingRuleInfo {
	return poolPontusTravelBookingRuleInfo.Get().(*PontusTravelBookingRuleInfo)
}

// ReleasePontusTravelBookingRuleInfo 释放PontusTravelBookingRuleInfo
func ReleasePontusTravelBookingRuleInfo(v *PontusTravelBookingRuleInfo) {
	v.RuleList = v.RuleList[:0]
	v.RuleDesc = ""
	v.RuleType = ""
	poolPontusTravelBookingRuleInfo.Put(v)
}
