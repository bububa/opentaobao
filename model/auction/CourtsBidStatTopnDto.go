package auction

// CourtsBidStatTopnDto 
type CourtsBidStatTopnDto struct {
    // 排名
    Rank   int64 `json:"rank,omitempty" xml:"rank,omitempty"`
    // 法院名称
    CourtName   string `json:"court_name,omitempty" xml:"court_name,omitempty"`
    // 成交价（成交标的）
    HammerPrice   int64 `json:"hammer_price,omitempty" xml:"hammer_price,omitempty"`
    // 发拍件数（去重）
    PublishCountDist   int64 `json:"publish_count_dist,omitempty" xml:"publish_count_dist,omitempty"`
}
