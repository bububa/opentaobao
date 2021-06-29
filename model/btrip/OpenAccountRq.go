package btrip

// OpenAccountRq 
type OpenAccountRq struct {
    // 对账单月份，不传取最新对账单
    BillMonth   string `json:"bill_month,omitempty" xml:"bill_month,omitempty"`
    // 第三方企业id
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 商旅开放平台传2
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
}
