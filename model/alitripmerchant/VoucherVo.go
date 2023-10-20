package alitripmerchant

import (
	"sync"
)

// VoucherVo 结构体
type VoucherVo struct {
	// 加价规则
	MarkUpRuleList []string `json:"mark_up_rule_list,omitempty" xml:"mark_up_rule_list>string,omitempty"`
	// 加价信息
	MarkUpInfo string `json:"mark_up_info,omitempty" xml:"mark_up_info,omitempty"`
	// 加价金额
	MarkUpAmount int64 `json:"mark_up_amount,omitempty" xml:"mark_up_amount,omitempty"`
	// 是否加价
	IsMarkUp bool `json:"is_mark_up,omitempty" xml:"is_mark_up,omitempty"`
}

var poolVoucherVo = sync.Pool{
	New: func() any {
		return new(VoucherVo)
	},
}

// GetVoucherVo() 从对象池中获取VoucherVo
func GetVoucherVo() *VoucherVo {
	return poolVoucherVo.Get().(*VoucherVo)
}

// ReleaseVoucherVo 释放VoucherVo
func ReleaseVoucherVo(v *VoucherVo) {
	v.MarkUpRuleList = v.MarkUpRuleList[:0]
	v.MarkUpInfo = ""
	v.MarkUpAmount = 0
	v.IsMarkUp = false
	poolVoucherVo.Put(v)
}
