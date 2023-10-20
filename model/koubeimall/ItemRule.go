package koubeimall

import (
	"sync"
)

// ItemRule 结构体
type ItemRule struct {
	// 规则列表
	RuleList []string `json:"rule_list,omitempty" xml:"rule_list>string,omitempty"`
}

var poolItemRule = sync.Pool{
	New: func() any {
		return new(ItemRule)
	},
}

// GetItemRule() 从对象池中获取ItemRule
func GetItemRule() *ItemRule {
	return poolItemRule.Get().(*ItemRule)
}

// ReleaseItemRule 释放ItemRule
func ReleaseItemRule(v *ItemRule) {
	v.RuleList = v.RuleList[:0]
	poolItemRule.Put(v)
}
