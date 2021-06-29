package scbp

// KeywordRankPriceDTO 
type KeywordRankPriceDTO struct {
    // 关键词
    Keyword   string `json:"keyword,omitempty" xml:"keyword,omitempty"`
    // 排价结果
    RankPriceList   []string `json:"rank_price_list,omitempty" xml:"rank_price_list>string,omitempty"`
}
