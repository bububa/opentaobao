package scbp

// KeywordRankPriceDto 
type KeywordRankPriceDto struct {

    // 关键词的id
    Keyword   string `json:"keyword,omitempty"`

    // 公司ID
    CompanyId   int64 `json:"company_id,omitempty"`

    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 价格列表(计划)(不建议直接使用)
    PriceList   []Number `json:"price_list,omitempty"`

    // 价格列表(计划)(元)(低价处理后结果)
    PriceArray   []String `json:"price_array,omitempty"`

    // 价格列表(客户)(不建议直接使用)
    CustPriceList   []Number `json:"cust_price_list,omitempty"`

    // 价格列表(客户)(元)(低价处理后结果)
    CustPriceArray   []String `json:"cust_price_array,omitempty"`

}
