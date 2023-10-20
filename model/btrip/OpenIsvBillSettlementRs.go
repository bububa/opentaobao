package btrip

import (
	"sync"
)

// OpenIsvBillSettlementRs 结构体
type OpenIsvBillSettlementRs struct {
	// 数据集合
	DataList []OpenIsvBillSettlementBtripTrainRs `json:"data_list,omitempty" xml:"data_list>open_isv_bill_settlement_btrip_train_rs,omitempty"`
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 记账更新开始时间
	PeriodStart string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	// 记账更新结束时间
	PeriodEnd string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// 总数据量
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 类目，枚举详见语雀
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
}

var poolOpenIsvBillSettlementRs = sync.Pool{
	New: func() any {
		return new(OpenIsvBillSettlementRs)
	},
}

// GetOpenIsvBillSettlementRs() 从对象池中获取OpenIsvBillSettlementRs
func GetOpenIsvBillSettlementRs() *OpenIsvBillSettlementRs {
	return poolOpenIsvBillSettlementRs.Get().(*OpenIsvBillSettlementRs)
}

// ReleaseOpenIsvBillSettlementRs 释放OpenIsvBillSettlementRs
func ReleaseOpenIsvBillSettlementRs(v *OpenIsvBillSettlementRs) {
	v.DataList = v.DataList[:0]
	v.CorpId = ""
	v.PeriodStart = ""
	v.PeriodEnd = ""
	v.TotalNum = 0
	v.Category = 0
	poolOpenIsvBillSettlementRs.Put(v)
}
