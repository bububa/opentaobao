package auction

// CourtsBidStatMonthlyList 
/* model for simplify = false
type CourtsBidStatMonthlyList struct {

    // 法院按月维度标的统计汇总
    
    CourtsBidStatSum  *struct {
        CourtsBidStatSum  *CourtsBidStatSum `json:"courts_bid_stat_sum,omitempty"`
    } `json:"courts_bid_stat_sum,omitempty"`
    

    // 时间区间(月份)
    
    Period   string `json:"period,omitempty"`
    

}
*/

// CourtsBidStatMonthlyList 
type CourtsBidStatMonthlyList struct {

    // 法院按月维度标的统计汇总
    CourtsBidStatSum   *CourtsBidStatSum `json:"courts_bid_stat_sum,omitempty"`

    // 时间区间(月份)
    Period   string `json:"period,omitempty"`

}
