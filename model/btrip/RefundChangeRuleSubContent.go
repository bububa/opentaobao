package btrip

import (
	"sync"
)

// RefundChangeRuleSubContent 结构体
type RefundChangeRuleSubContent struct {
}

var poolRefundChangeRuleSubContent = sync.Pool{
	New: func() any {
		return new(RefundChangeRuleSubContent)
	},
}

// GetRefundChangeRuleSubContent() 从对象池中获取RefundChangeRuleSubContent
func GetRefundChangeRuleSubContent() *RefundChangeRuleSubContent {
	return poolRefundChangeRuleSubContent.Get().(*RefundChangeRuleSubContent)
}

// ReleaseRefundChangeRuleSubContent 释放RefundChangeRuleSubContent
func ReleaseRefundChangeRuleSubContent(v *RefundChangeRuleSubContent) {
	poolRefundChangeRuleSubContent.Put(v)
}
