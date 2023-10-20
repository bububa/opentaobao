package txcs

import (
	"sync"
)

// BizExtAttr 结构体
type BizExtAttr struct {
	// 汇总编码
	SummaryId string `json:"summary_id,omitempty" xml:"summary_id,omitempty"`
	// 账单创建时间
	BillCreateTime string `json:"bill_create_time,omitempty" xml:"bill_create_time,omitempty"`
}

var poolBizExtAttr = sync.Pool{
	New: func() any {
		return new(BizExtAttr)
	},
}

// GetBizExtAttr() 从对象池中获取BizExtAttr
func GetBizExtAttr() *BizExtAttr {
	return poolBizExtAttr.Get().(*BizExtAttr)
}

// ReleaseBizExtAttr 释放BizExtAttr
func ReleaseBizExtAttr(v *BizExtAttr) {
	v.SummaryId = ""
	v.BillCreateTime = ""
	poolBizExtAttr.Put(v)
}
