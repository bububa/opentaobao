package moscm

import (
	"sync"
)

// RmaCriteria 结构体
type RmaCriteria struct {
	// 退换货单据号
	RmaNumbers []string `json:"rma_numbers,omitempty" xml:"rma_numbers>string,omitempty"`
	// 单据状态，可选值：已创建(&#34;CREATED&#34;),已收货同意退款(&#34;INBOUND&#34;),已收货不同意退款(&#34;NOTINBOUND&#34;),已退款(&#34;REFUNDED&#34;),已完成(&#34;COMPLETED),已作废(&#34;Obsolete&#34;)
	Status []string `json:"status,omitempty" xml:"status>string,omitempty"`
	// 订单号
	OrderNumber string `json:"order_number,omitempty" xml:"order_number,omitempty"`
	// 单据创建时间范围：结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 单据创建时间范围：开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 单据类型，可选值：退货(&#34;GOODRETURN&#34;),换货(&#34;GOODEXCHANGE&#34;),仅退款（&#34;RETURN&#34;）
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 供应商专柜Id
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 银泰专柜Id
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
}

var poolRmaCriteria = sync.Pool{
	New: func() any {
		return new(RmaCriteria)
	},
}

// GetRmaCriteria() 从对象池中获取RmaCriteria
func GetRmaCriteria() *RmaCriteria {
	return poolRmaCriteria.Get().(*RmaCriteria)
}

// ReleaseRmaCriteria 释放RmaCriteria
func ReleaseRmaCriteria(v *RmaCriteria) {
	v.RmaNumbers = v.RmaNumbers[:0]
	v.Status = v.Status[:0]
	v.OrderNumber = ""
	v.EndDate = ""
	v.StartDate = ""
	v.Type = ""
	v.OutCounterId = ""
	v.CounterId = ""
	poolRmaCriteria.Put(v)
}
