package tmallservice

import (
	"sync"
)

// FeeList 结构体
type FeeList struct {
	// 费用来源单号，仅增加费用、退款有值
	SrcOrderId string `json:"src_order_id,omitempty" xml:"src_order_id,omitempty"`
	// 费用描述
	FeeTitle string `json:"fee_title,omitempty" xml:"fee_title,omitempty"`
	// 费用名称
	FeeName string `json:"fee_name,omitempty" xml:"fee_name,omitempty"`
	// 费用产生时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 费用金额
	FeeAmount float64 `json:"fee_amount,omitempty" xml:"fee_amount,omitempty"`
}

var poolFeeList = sync.Pool{
	New: func() any {
		return new(FeeList)
	},
}

// GetFeeList() 从对象池中获取FeeList
func GetFeeList() *FeeList {
	return poolFeeList.Get().(*FeeList)
}

// ReleaseFeeList 释放FeeList
func ReleaseFeeList(v *FeeList) {
	v.SrcOrderId = ""
	v.FeeTitle = ""
	v.FeeName = ""
	v.GmtCreate = ""
	v.FeeAmount = 0
	poolFeeList.Put(v)
}
