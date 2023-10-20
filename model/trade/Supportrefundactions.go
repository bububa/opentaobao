package trade

import (
	"sync"
)

// Supportrefundactions 结构体
type Supportrefundactions struct {
	// 退款退货操作的描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 退款退货操作按钮的建议描述，仅仅是一个建议，ISV可以自己定义
	DefaultLabel string `json:"default_label,omitempty" xml:"default_label,omitempty"`
	// 退款退货操作的Code，由系统定义，目前支持的方式有：refundFeeOnly(仅退款)，refundFeeWithGoods(退货退款),swithGoods(换货)
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 一个纠纷单可能已经在处理流程中，比如退款退货操作，买家已经提交申请，卖家正在审核中，则该字段是true
	InProcess string `json:"in_process,omitempty" xml:"in_process,omitempty"`
}

var poolSupportrefundactions = sync.Pool{
	New: func() any {
		return new(Supportrefundactions)
	},
}

// GetSupportrefundactions() 从对象池中获取Supportrefundactions
func GetSupportrefundactions() *Supportrefundactions {
	return poolSupportrefundactions.Get().(*Supportrefundactions)
}

// ReleaseSupportrefundactions 释放Supportrefundactions
func ReleaseSupportrefundactions(v *Supportrefundactions) {
	v.Desc = ""
	v.DefaultLabel = ""
	v.Key = ""
	v.InProcess = ""
	poolSupportrefundactions.Put(v)
}
