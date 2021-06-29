package scbp

// KeywordRankPriceDto 
type KeywordRankPriceDto struct {
    // 关键词的id
    Keyword   string `json:"keyword,omitempty" xml:"keyword,omitempty"`
    // 公司ID
    CompanyId   int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 价格列表(计划)(不建议直接使用)
    PriceList   []int64 `json:"price_list,omitempty" xml:"price_list>int64,omitempty"`
    // 价格列表(计划)(元)(低价处理后结果)
    PriceArray   []string `json:"price_array,omitempty" xml:"price_array>string,omitempty"`
    // 价格列表(客户)(不建议直接使用)
    CustPriceList   []int64 `json:"cust_price_list,omitempty" xml:"cust_price_list>int64,omitempty"`
    // 价格列表(客户)(元)(低价处理后结果)
    CustPriceArray   []string `json:"cust_price_array,omitempty" xml:"cust_price_array>string,omitempty"`
}
