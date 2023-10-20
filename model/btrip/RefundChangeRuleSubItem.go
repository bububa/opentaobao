package btrip

import (
	"sync"
)

// RefundChangeRuleSubItem 结构体
type RefundChangeRuleSubItem struct {
	// 退款子内容
	RefundSubContents []RefundChangeRuleSubContent `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents>refund_change_rule_sub_content,omitempty"`
	// PTC
	Ptc string `json:"ptc,omitempty" xml:"ptc,omitempty"`
	// 退改签规则的类型
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 是否结构体
	IsStruct bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
}

var poolRefundChangeRuleSubItem = sync.Pool{
	New: func() any {
		return new(RefundChangeRuleSubItem)
	},
}

// GetRefundChangeRuleSubItem() 从对象池中获取RefundChangeRuleSubItem
func GetRefundChangeRuleSubItem() *RefundChangeRuleSubItem {
	return poolRefundChangeRuleSubItem.Get().(*RefundChangeRuleSubItem)
}

// ReleaseRefundChangeRuleSubItem 释放RefundChangeRuleSubItem
func ReleaseRefundChangeRuleSubItem(v *RefundChangeRuleSubItem) {
	v.RefundSubContents = v.RefundSubContents[:0]
	v.Ptc = ""
	v.Title = ""
	v.IsStruct = false
	poolRefundChangeRuleSubItem.Put(v)
}
