package scbp

// KeywordRankPriceDTO 
type KeywordRankPriceDTO struct {

    // 关键词
    Keyword   string `json:"keyword,omitempty"`

    // 排价结果
    RankPriceList   []String `json:"rank_price_list,omitempty"`

}
