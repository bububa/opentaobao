package btrip

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
