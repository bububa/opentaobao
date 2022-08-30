package btrip

// OpenIsvBillSettlementSearchRq 结构体
type OpenIsvBillSettlementSearchRq struct {
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 记账更新开始日期
	PeriodStart string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	// 记账更新结束日期
	PeriodEnd string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// 每页数据量，默认100，最高100
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 1、老版本2、isv对外版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 页数，从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 类目：机酒火车 1：机票； 2：酒店； 4：用车 6：商旅火车票
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
}
