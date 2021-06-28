package traderate

// GetMixRateListResult 
/* model for simplify = false
type GetMixRateListResult struct {

    // 是否下一页
    
    HasNextPage   int64 `json:"has_next_page,omitempty"`
    

    // 评价统计信息
    
    ItemStatistic  *struct {
        ItemStatisticVo  *ItemStatisticVo `json:"item_statistic_vo,omitempty"`
    } `json:"item_statistic,omitempty"`
    

    // 评价明细信息
    
    MixRates  struct {
        MixRateVo  []MixRateVo `json:"mix_rate_vo,omitempty"`
    } `json:"mix_rates,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 总数量
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

}
*/

// GetMixRateListResult 
type GetMixRateListResult struct {

    // 是否下一页
    HasNextPage   int64 `json:"has_next_page,omitempty"`

    // 评价统计信息
    ItemStatistic   *ItemStatisticVo `json:"item_statistic,omitempty"`

    // 评价明细信息
    MixRates   []MixRateVo `json:"mix_rates,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 总数量
    TotalNum   int64 `json:"total_num,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

}
