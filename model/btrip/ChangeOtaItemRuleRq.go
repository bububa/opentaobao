package btrip

import (
	"sync"
)

// ChangeOtaItemRuleRq 结构体
type ChangeOtaItemRuleRq struct {
	// 退票规则
	RefundDetails []ChangeOtaItemRuleRq `json:"refund_details,omitempty" xml:"refund_details>change_ota_item_rule_rq,omitempty"`
	// 改签规则
	ChangeDetails []RefundDetailShowVo `json:"change_details,omitempty" xml:"change_details>refund_detail_show_vo,omitempty"`
	// 行李额
	BaggageDetails []RefundDetailShowVo `json:"baggage_details,omitempty" xml:"baggage_details>refund_detail_show_vo,omitempty"`
	// 退改表达内容
	RefundSubItems []RefundSubItem `json:"refund_sub_items,omitempty" xml:"refund_sub_items>refund_sub_item,omitempty"`
	// 纯文字段落
	ExtraContents []RefundSubItem `json:"extra_contents,omitempty" xml:"extra_contents>refund_sub_item,omitempty"`
	// 表格标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 表格首行内容
	TableHead string `json:"table_head,omitempty" xml:"table_head,omitempty"`
	// 内容类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 排序使用
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
}

var poolChangeOtaItemRuleRq = sync.Pool{
	New: func() any {
		return new(ChangeOtaItemRuleRq)
	},
}

// GetChangeOtaItemRuleRq() 从对象池中获取ChangeOtaItemRuleRq
func GetChangeOtaItemRuleRq() *ChangeOtaItemRuleRq {
	return poolChangeOtaItemRuleRq.Get().(*ChangeOtaItemRuleRq)
}

// ReleaseChangeOtaItemRuleRq 释放ChangeOtaItemRuleRq
func ReleaseChangeOtaItemRuleRq(v *ChangeOtaItemRuleRq) {
	v.RefundDetails = v.RefundDetails[:0]
	v.ChangeDetails = v.ChangeDetails[:0]
	v.BaggageDetails = v.BaggageDetails[:0]
	v.RefundSubItems = v.RefundSubItems[:0]
	v.ExtraContents = v.ExtraContents[:0]
	v.Title = ""
	v.TableHead = ""
	v.Type = 0
	v.Index = 0
	poolChangeOtaItemRuleRq.Put(v)
}
