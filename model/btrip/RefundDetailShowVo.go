package btrip

import (
	"sync"
)

// RefundDetailShowVo 结构体
type RefundDetailShowVo struct {
	// 退改表达内容
	RefundSubItems []RefundSubItem `json:"refund_sub_items,omitempty" xml:"refund_sub_items>refund_sub_item,omitempty"`
	// 纯文字段落
	ExtraContents []RefundSubItem `json:"extra_contents,omitempty" xml:"extra_contents>refund_sub_item,omitempty"`
	// 行李表达内容
	BaggageSubItems []BaggageSubItem `json:"baggage_sub_items,omitempty" xml:"baggage_sub_items>baggage_sub_item,omitempty"`
	// 表格标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 表格首行内容
	TableHead string `json:"table_head,omitempty" xml:"table_head,omitempty"`
	// 内容类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 排序使用
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
	// 提示
	Tips *Tips `json:"tips,omitempty" xml:"tips,omitempty"`
}

var poolRefundDetailShowVo = sync.Pool{
	New: func() any {
		return new(RefundDetailShowVo)
	},
}

// GetRefundDetailShowVo() 从对象池中获取RefundDetailShowVo
func GetRefundDetailShowVo() *RefundDetailShowVo {
	return poolRefundDetailShowVo.Get().(*RefundDetailShowVo)
}

// ReleaseRefundDetailShowVo 释放RefundDetailShowVo
func ReleaseRefundDetailShowVo(v *RefundDetailShowVo) {
	v.RefundSubItems = v.RefundSubItems[:0]
	v.ExtraContents = v.ExtraContents[:0]
	v.BaggageSubItems = v.BaggageSubItems[:0]
	v.Title = ""
	v.TableHead = ""
	v.Type = 0
	v.Index = 0
	v.Tips = nil
	poolRefundDetailShowVo.Put(v)
}
