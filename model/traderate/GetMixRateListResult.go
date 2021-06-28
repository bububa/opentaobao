package traderate

// GetMixRateListResult 
type GetMixRateListResult struct {

    // 是否下一页
    
    HasNextPage   int64 `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
    

    // 评价统计信息
    
    ItemStatistic   *ItemStatisticVo `json:"item_statistic,omitempty" xml:"item_statistic,omitempty"`
    

    // 评价明细信息
    
    MixRates   []MixRateVo `json:"mix_rates,omitempty" xml:"mix_rates,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 总数量
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

}
